package postgresql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/forbole/juno/v5/types"
)

// HasBlock implements database.Database
func (db *Database) HasBlock(height int64) (bool, error) {
	var res bool
	err := db.SQL.QueryRow(`SELECT EXISTS(SELECT 1 FROM blocks WHERE height = $1)`, height).Scan(&res)
	return res, err
}

// GetLastBlockHeight returns the last block height stored inside the database
func (db *Database) GetLastBlockHeight() (int64, error) {
	stmt := `SELECT height FROM blocks ORDER BY height DESC LIMIT 1`

	var height int64
	err := db.SQL.QueryRow(stmt).Scan(&height)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("error while getting last block height, error: %s", err)
	}

	return height, nil
}

// GetMissingHeights returns a slice of missing block heights between startHeight and endHeight
func (db *Database) GetMissingHeights(startHeight, endHeight int64) []int64 {
	var result []int64
	stmt := `SELECT generate_series($1::int,$2::int) EXCEPT SELECT height FROM blocks ORDER BY 1`

	err := db.SQL.Select(&result, stmt, startHeight, endHeight)
	if err != nil {
		return nil
	}

	return result
}

// SaveBlock implements database.Database
func (db *Database) SaveBlock(block *types.Block) error {
	sqlStatement := `
INSERT INTO blocks (height, hash, num_txs, total_gas, proposer_address, timestamp)
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING`

	proposerAddress := sql.NullString{Valid: len(block.ProposerAddress) != 0, String: block.ProposerAddress}
	_, err := db.SQL.Exec(sqlStatement,
		block.Height,
		block.Hash,
		block.TxNum,
		block.TotalGas,
		proposerAddress,
		block.Timestamp,
	)
	return err
}

// GetTotalBlocks implements database.Database
func (db *Database) GetTotalBlocks() int64 {
	var blockCount int64
	err := db.SQL.QueryRow(`SELECT count(*) FROM blocks`).Scan(&blockCount)
	if err != nil {
		return 0
	}

	return blockCount
}
