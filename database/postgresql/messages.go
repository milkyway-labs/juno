package postgresql

import (
	"github.com/jmoiron/sqlx"

	"github.com/forbole/juno/v5/types"
)

type involvedAddressRow struct {
	UserAddress     string `db:"user_address"`
	MessageIndex    int    `db:"message_index"`
	TransactionHash string `db:"transaction_hash"`
	PartitionID     int64  `db:"partition_id"`
}

// SaveMessage implements database.Database
func (db *Database) SaveMessage(msg *types.Message) error {
	return db.SaveMessages([]*types.Message{msg})
}

// SaveMessages allows to store multiple messages at the same time
func (db *Database) SaveMessages(messages []*types.Message) error {
	tx, err := db.SQL.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = db.saveMessagesWithTx(tx, messages)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// saveMessagesWithTx stores the given messages inside the database.
func (db *Database) saveMessagesWithTx(tx *sqlx.Tx, messages []*types.Message) error {
	for _, msg := range messages {
		partitionID := msg.Height / db.partitionSize
		err := db.createPartitionIfNotExistsWithTx(tx, "messages", partitionID)
		if err != nil {
			return err
		}

		err = db.saveMessageInsidePartition(tx, msg, partitionID)
		if err != nil {
			return err
		}
	}

	return nil
}

// saveMessageInsidePartition stores the given message inside the partition having the provided id.
// All the operations are done inside the same transaction.
// It is responsibility of the caller to commit the transaction or rollback.
func (db *Database) saveMessageInsidePartition(tx *sqlx.Tx, msg *types.Message, partitionID int64) error {
	// Store the message
	stmt := `
INSERT INTO messages (index, type, value, transaction_hash, partition_id) 
VALUES ($1, $2, $3, $4, $5) 
ON CONFLICT ON CONSTRAINT unique_message_per_tx DO UPDATE 
	SET  type = excluded.type,
	     value = excluded.value`
	_, err := tx.Exec(stmt, msg.Index, msg.Type, msg.Value, msg.TxHash, partitionID)
	if err != nil {
		return err
	}

	// Store the involved addresses
	rows := make([]involvedAddressRow, len(msg.Addresses))
	for i, address := range msg.Addresses {
		rows[i] = involvedAddressRow{
			UserAddress:     address,
			MessageIndex:    msg.Index,
			TransactionHash: msg.TxHash,
			PartitionID:     partitionID,
		}
	}

	if len(rows) == 0 {
		return nil
	}

	// Create the partition if it does not exist
	err = db.createPartitionIfNotExistsWithTx(tx, "message_involved_accounts", partitionID)
	if err != nil {
		return err
	}

	// Store the involved accounts
	stmt = `
INSERT INTO message_involved_accounts (user_address, message_index, transaction_hash, partition_id)
VALUES (:user_address, :message_index, :transaction_hash, :partition_id)`
	_, err = tx.NamedExec(stmt, rows)

	return err
}
