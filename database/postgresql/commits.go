package postgresql

import (
	"time"

	"github.com/forbole/juno/v5/types"
	"github.com/forbole/juno/v5/utils"
)

type preCommitRow struct {
	ValidatorAddress string    `db:"validator_address"`
	Height           int64     `db:"height"`
	Timestamp        time.Time `db:"timestamp"`
	VotingPower      int64     `db:"voting_power"`
	ProposerPriority int64     `db:"proposer_priority"`
}

func convertCommitSigToRow(row *types.CommitSig) preCommitRow {
	return preCommitRow{
		ValidatorAddress: row.ValidatorAddress,
		Height:           row.Height,
		Timestamp:        row.Timestamp,
		VotingPower:      row.VotingPower,
		ProposerPriority: row.ProposerPriority,
	}
}

// SaveCommitSignatures implements database.Database
func (db *Database) SaveCommitSignatures(signatures []*types.CommitSig) error {
	rows := utils.Map(signatures, convertCommitSigToRow)
	if len(rows) == 0 {
		return nil
	}

	stmt := `
INSERT INTO pre_commits (validator_address, height, timestamp, voting_power, proposer_priority) 
VALUES (:validator_address, :height, :timestamp, :voting_power, :proposer_priority)
ON CONFLICT DO NOTHING`
	_, err := db.SQL.NamedExec(stmt, rows)
	return err
}
