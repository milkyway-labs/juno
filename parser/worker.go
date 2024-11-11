package parser

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/forbole/juno/v5/prometheus"
	"github.com/forbole/juno/v5/utils"

	"github.com/forbole/juno/v5/database"
	"github.com/forbole/juno/v5/logging"
	nodeutils "github.com/forbole/juno/v5/node/utils"
	"github.com/forbole/juno/v5/types/config"

	"github.com/forbole/juno/v5/modules"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"

	sdk "github.com/forbole/juno/v5/cosmos-sdk/types"

	codectypes "github.com/forbole/juno/v5/cosmos-sdk/codec/types"
	"github.com/forbole/juno/v5/node"
	"github.com/forbole/juno/v5/types"
)

// Worker defines a job consumer that is responsible for getting and
// aggregating block and associated data and exporting it to a database.
type Worker struct {
	index int

	cfg config.Config

	queue types.HeightQueue

	modules []modules.Module

	node   node.Node
	db     database.Database
	logger logging.Logger
}

// NewWorker allows to create a new Worker implementation.
func NewWorker(ctx *Context, queue types.HeightQueue, index int) Worker {
	return Worker{
		index:   index,
		cfg:     ctx.Config,
		node:    ctx.Node,
		queue:   queue,
		db:      ctx.Database,
		modules: ctx.Modules,
		logger:  ctx.Logger,
	}
}

// shouldReEnqueueWhenFailed returns true if the worker should re-enqueue a block when
// the parsing of its parts fails inside modules
func (w Worker) shouldReEnqueueWhenFailed() bool {
	return w.cfg.Parser.ReEnqueueWhenFailed
}

// Start starts a worker by listening for new jobs (block heights) from the
// given worker queue. Any failed job is logged and re-enqueued.
func (w Worker) Start() {
	prometheus.WorkersCount.Inc()

	for i := range w.queue {
		// Make sure we did not reach the max retries yet
		if i.HasReachedMaxRetries(w.cfg.Parser.GetMaxRetries()) {
			w.logger.Error("failed to process block", "height", i, "err", i.LatestError)
			continue
		}

		// Process the block
		err := w.ProcessIfNotExists(i.Height)
		if err != nil {
			go func() {
				// Signal that an error occurred while processing this block
				prometheus.SignalBlockError(i.Height)

				// Build the block with the updated retry count and log the error
				newBlock := i.IncrementRetryCount(err)
				w.logger.Debug("re-enqueuing failed block", "height", i.Height, "err", err, "count", newBlock.RetryCount)

				// Sleep for the proper time and re-enqueue the block
				time.Sleep(w.cfg.GetAvgBlockTime() * time.Duration(newBlock.RetryCount))
				w.queue <- newBlock
			}()
		}

		prometheus.LatestIndexedHeightByWorker.WithLabelValues(fmt.Sprintf("%d", w.index)).Set(float64(i.Height))
	}
}

// ProcessIfNotExists defines the job consumer workflow. It will fetch a block for a given
// height and associated metadata and export it to a database if it does not exist yet. It returns an
// error if any export process fails.
func (w Worker) ProcessIfNotExists(height int64) error {
	exists, err := w.db.HasBlock(height)
	if err != nil {
		return fmt.Errorf("error while searching for block: %s", err)
	}

	// If the block already exists and the height is not included in the reparse range, skip it
	if exists {
		if w.cfg.Parser.ReparseRange != nil && !w.cfg.Parser.ReparseRange.Includes(height) {
			w.logger.Debug("skipping already exported block", "height", height)
			return nil
		}

		w.logger.Debug("re-parsing block", "height", height)
	}

	return w.Process(height)
}

// Process fetches  a block for a given height and associated metadata and export it to a database.
// It returns an error if any export process fails.
func (w Worker) Process(height int64) error {
	if height == 0 {
		cfg := w.cfg.Parser

		genesisDoc, genesisState, err := nodeutils.GetGenesisDocAndState(cfg.GenesisFilePath, w.node)
		if err != nil {
			return fmt.Errorf("failed to get genesis: %s", err)
		}

		return w.HandleGenesis(genesisDoc, genesisState)
	}

	w.logger.Debug("processing block", "height", height)

	block, err := w.node.Block(height)
	if err != nil {
		return fmt.Errorf("failed to get block from node: %s", err)
	}

	events, err := w.node.BlockResults(height)
	if err != nil {
		return fmt.Errorf("failed to get block results from node: %s", err)
	}

	txs, err := w.node.Txs(block)
	if err != nil {
		return fmt.Errorf("failed to get transactions for block: %s", err)
	}

	vals, err := w.node.Validators(height)
	if err != nil {
		return fmt.Errorf("failed to get validators for block: %s", err)
	}

	return w.ExportBlock(block, events, txs, vals)
}

// ProcessTransactions fetches transactions for a given height and stores them into the database.
// It returns an error if the export process fails.
func (w Worker) ProcessTransactions(height int64) error {
	block, err := w.node.Block(height)
	if err != nil {
		return fmt.Errorf("failed to get block from node: %s", err)
	}

	txs, err := w.node.Txs(block)
	if err != nil {
		return fmt.Errorf("failed to get transactions for block: %s", err)
	}

	return w.ExportTxs(txs)
}

// HandleGenesis accepts a GenesisDoc and calls all the registered genesis handlers
// in the order in which they have been registered.
func (w Worker) HandleGenesis(genesisDoc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	// Call the genesis handlers
	for _, module := range w.modules {
		if genesisModule, ok := module.(modules.GenesisModule); ok {
			err := genesisModule.HandleGenesis(genesisDoc, appState)
			if err != nil {
				if w.shouldReEnqueueWhenFailed() {
					return err
				}

				w.logger.GenesisError(module, err)
			}
		}
	}

	return nil
}

// SaveValidators persists a list of Tendermint validators with an address and a
// consensus public key. An error is returned if the public key cannot be Bech32
// encoded or if the DB write fails.
func (w Worker) SaveValidators(vals []*tmtypes.Validator) error {
	validators := make([]*types.Validator, len(vals))
	for index, val := range vals {
		consAddr := sdk.NewConsAddress(val.Address).String()

		consPubKey, err := utils.ConvertValidatorPubKeyToBech32String(val.PubKey)
		if err != nil {
			return fmt.Errorf("failed to convert validator public key for validators %s: %s", consAddr, err)
		}

		validators[index] = types.NewValidator(consAddr, consPubKey)
	}

	err := w.db.SaveValidators(validators)
	if err != nil {
		return fmt.Errorf("error while saving validators: %s", err)
	}

	return nil
}

// ExportBlock accepts a finalized block and a corresponding set of transactions
// and persists them to the database along with attributable metadata. An error
// is returned if the write fails.
func (w Worker) ExportBlock(
	b *tmctypes.ResultBlock, r *tmctypes.ResultBlockResults, txs []*types.Tx, vals *tmctypes.ResultValidators,
) error {
	// Save all validators
	err := w.SaveValidators(vals.Validators)
	if err != nil {
		return err
	}

	// Make sure the proposer exists
	proposerAddr := sdk.NewConsAddress(b.Block.ProposerAddress)
	val := findValidatorByAddr(proposerAddr.String(), vals)
	if val == nil {
		return fmt.Errorf("failed to find validator by proposer address %s: %s", proposerAddr.String(), err)
	}

	// Save the block
	err = w.db.SaveBlock(types.NewBlockFromTmBlock(b, sumGasTxs(txs)))
	if err != nil {
		return fmt.Errorf("failed to persist block: %s", err)
	}

	// Save the commits
	err = w.ExportCommit(b.Block.LastCommit, vals)
	if err != nil {
		return err
	}

	// Call the block handlers
	for _, module := range w.modules {
		if blockModule, ok := module.(modules.BlockModule); ok {
			err = blockModule.HandleBlock(b, r, txs, vals)
			if err != nil {
				if w.shouldReEnqueueWhenFailed() {
					return err
				}

				w.logger.BlockError(module, b, err)
			}
		}
	}

	// Export the transactions
	return w.ExportTxs(txs)
}

// ExportCommit accepts a block commitment and a corresponding set of
// validators for the commitment and persists them to the database. An error is
// returned if any write fails or if there is any missing aggregated data.
func (w Worker) ExportCommit(commit *tmtypes.Commit, vals *tmctypes.ResultValidators) error {
	var signatures []*types.CommitSig
	for _, commitSig := range commit.Signatures {
		// Avoid empty commits
		if commitSig.Signature == nil {
			continue
		}

		valAddr := sdk.NewConsAddress(commitSig.ValidatorAddress)
		val := findValidatorByAddr(valAddr.String(), vals)
		if val == nil {
			return fmt.Errorf("failed to find validator by commit validator address %s", valAddr.String())
		}

		signatures = append(signatures, types.NewCommitSig(
			utils.ConvertValidatorAddressToBech32String(commitSig.ValidatorAddress),
			val.VotingPower,
			val.ProposerPriority,
			commit.Height,
			commitSig.Timestamp,
		))
	}

	err := w.db.SaveCommitSignatures(signatures)
	if err != nil {
		return fmt.Errorf("error while saving commit signatures: %s", err)
	}

	return nil
}

// saveTx accepts the transaction and persists it inside the database.
// An error is returned if the write fails.
func (w Worker) saveTx(tx *types.Tx) error {
	err := w.db.SaveTx(tx)
	if err != nil {
		return fmt.Errorf("failed to handle transaction with hash %s: %s", tx.TxHash, err)
	}
	return nil
}

// handleTx accepts the transaction and calls the tx handlers.
func (w Worker) handleTx(tx *types.Tx) error {
	// Call the tx handlers
	for _, module := range w.modules {
		if transactionModule, ok := module.(modules.TransactionModule); ok {
			err := transactionModule.HandleTx(tx)
			if err != nil {
				if w.shouldReEnqueueWhenFailed() {
					return err
				}

				w.logger.TxError(module, tx, err)
			}
		}
	}

	return nil
}

// handleMessage accepts the transaction and handles messages contained
// inside the transaction.
func (w Worker) handleMessage(index int, msg *codectypes.Any, tx *types.Tx) error {
	// Allow modules to handle the message
	for _, module := range w.modules {
		if messageModule, ok := module.(modules.MessageModule); ok {
			err := messageModule.HandleMsg(index, msg, tx)
			if err != nil {
				if w.shouldReEnqueueWhenFailed() {
					w.logger.MsgError(module, tx, msg, err)
				}
			}
		}
	}

	return nil
}

// ExportTxs accepts a slice of transactions and persists then inside the database.
// An error is returned if the write fails.
func (w Worker) ExportTxs(txs []*types.Tx) error {
	for _, tx := range txs {
		// Save the transaction
		err := w.saveTx(tx)
		if err != nil {
			return fmt.Errorf("error while storing txs: %s", err)
		}

		// Call the transactions handlers
		err = w.handleTx(tx)
		if err != nil {
			return err
		}

		// Call the message handlers
		for i, sdkMsg := range tx.Body.Messages {
			err = w.handleMessage(i, sdkMsg, tx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
