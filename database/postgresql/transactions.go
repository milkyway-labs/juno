package postgresql

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/lib/pq"

	"github.com/forbole/juno/v5/types"
)

// SaveTx implements database.Database
func (db *Database) SaveTx(tx *types.Tx) error {
	// Create the partition
	partitionID := tx.Height / db.partitionSize
	err := db.CreatePartitionIfNotExists("transactions", partitionID)
	if err != nil {
		return err
	}

	// Store the transaction data
	return db.saveTxInsidePartition(tx, partitionID)
}

// saveTxInsidePartition stores the given transaction inside the partition having the given id
func (db *Database) saveTxInsidePartition(tx *types.Tx, partitionID int64) error {
	// Do not store transactions that we do not want to store
	if !db.ShouldStoreTransaction(tx) {
		return nil
	}

	// Start a transaction
	dbTx, err := db.SQL.Beginx()
	if err != nil {
		return err
	}
	defer dbTx.Rollback()

	// Store the transaction
	stmt := `
INSERT INTO transactions 
(hash, height, success, memo, signatures, signer_infos, fee, gas_wanted, gas_used, raw_log, logs, partition_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
ON CONFLICT (hash, partition_id) DO UPDATE 
	SET height = excluded.height, 
		success = excluded.success, 
		memo = excluded.memo, 
		signatures = excluded.signatures, 
		signer_infos = excluded.signer_infos,
		fee = excluded.fee, 
		gas_wanted = excluded.gas_wanted, 
		gas_used = excluded.gas_used,
		raw_log = excluded.raw_log, 
		logs = excluded.logs`

	var sigs = make([]string, len(tx.Signatures))
	for index, sig := range tx.Signatures {
		sigs[index] = base64.StdEncoding.EncodeToString(sig)
	}

	feeBz, err := db.Cdc.MarshalJSON(tx.AuthInfo.Fee)
	if err != nil {
		return fmt.Errorf("failed to JSON encode tx fee: %s", err)
	}

	var sigInfos = make([]string, len(tx.AuthInfo.SignerInfos))
	for index, info := range tx.AuthInfo.SignerInfos {
		bz, err := db.Cdc.MarshalJSON(info)
		if err != nil {
			return err
		}
		sigInfos[index] = string(bz)
	}
	sigInfoBz := fmt.Sprintf("[%s]", strings.Join(sigInfos, ","))

	logsBz, err := db.Amino.MarshalJSON(tx.Logs)
	if err != nil {
		return err
	}

	_, err = dbTx.Exec(stmt,
		tx.TxHash,
		tx.Height,
		tx.Successful(),
		tx.Body.Memo,
		pq.Array(sigs),
		sigInfoBz,
		string(feeBz),
		tx.GasWanted,
		tx.GasUsed,
		tx.RawLog,
		string(logsBz),
		partitionID,
	)
	if err != nil {
		return err
	}

	// Commit the transaction
	return dbTx.Commit()
}
