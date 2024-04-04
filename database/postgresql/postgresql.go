package postgresql

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/jmoiron/sqlx"

	"github.com/forbole/juno/v5/logging"
	"github.com/forbole/juno/v5/types"
	"github.com/forbole/juno/v5/utils"

	"github.com/forbole/juno/v5/database"
)

// type check to ensure interface is properly implemented
var _ database.Database = &Database{}

// Database defines a wrapper around a SQL database and implements functionality
// for data aggregation and exporting.
type Database struct {
	partitionSize int64

	Cdc   codec.Codec
	Amino *codec.LegacyAmino

	SQL                  *sqlx.DB
	Logger               logging.Logger
	AccountAddressParser types.AccountAddressParser
}

// Builder creates a database connection with the given database connection info
// from config. It returns a database connection handle or an error if the
// connection fails.
func Builder(ctx *database.Context) (database.Database, error) {
	dbURI := utils.GetEnvOr(types.DatabaseURI, ctx.Cfg.URL)

	postgresDb, err := sqlx.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	return &Database{
		partitionSize: ctx.Cfg.GetPartitionSize(),

		Cdc:   ctx.EncodingConfig.Codec,
		Amino: ctx.EncodingConfig.Amino,

		SQL:                  postgresDb,
		Logger:               ctx.Logger,
		AccountAddressParser: ctx.AccountAddressParser,
	}, nil
}

// CreatePartitionIfNotExists creates a new partition having the given partition id if not existing
func (db *Database) CreatePartitionIfNotExists(table string, partitionID int64) error {
	// Create a transaction
	tx, err := db.SQL.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = db.createPartitionIfNotExistsWithTx(tx, table, partitionID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (db *Database) createPartitionIfNotExistsWithTx(tx *sqlx.Tx, table string, partitionID int64) error {
	stmt := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s PARTITION OF %s FOR VALUES IN (%d)",
		fmt.Sprintf("%s_%d", table, partitionID),
		table,
		partitionID,
	)
	_, err := tx.Exec(stmt)
	return err
}

// Close implements database.Database
func (db *Database) Close() {
	err := db.SQL.Close()
	if err != nil {
		db.Logger.Error("error while closing connection", "err", err)
	}
}
