package postgresql

// GetLastPruned implements database.PruningDb
func (db *Database) GetLastPruned() (int64, error) {
	var lastPrunedHeight int64
	err := db.SQL.QueryRow(`SELECT coalesce(MAX(last_pruned_height),0) FROM pruning LIMIT 1`).Scan(&lastPrunedHeight)
	return lastPrunedHeight, err
}

// StoreLastPruned implements database.PruningDb
func (db *Database) StoreLastPruned(height int64) error {
	_, err := db.SQL.Exec(`DELETE FROM pruning WHERE TRUE`)
	if err != nil {
		return err
	}

	_, err = db.SQL.Exec(`INSERT INTO pruning (last_pruned_height) VALUES ($1)`, height)
	return err
}

// Prune implements database.PruningDb
func (db *Database) Prune(height int64) error {
	_, err := db.SQL.Exec(`DELETE FROM pre_commits WHERE height = $1`, height)
	if err != nil {
		return err
	}

	_, err = db.SQL.Exec(`
DELETE FROM messages 
USING transactions 
WHERE messages.transaction_hash = transactions.hash AND transactions.height = $1
`, height)
	return err
}
