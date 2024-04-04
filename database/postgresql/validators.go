package postgresql

import (
	"github.com/forbole/juno/v5/types"
	"github.com/forbole/juno/v5/utils"
)

type validatorDbRow struct {
	ConsensusAddress string `db:"consensus_address"`
	ConsensusPubKey  string `db:"consensus_pubkey"`
}

func convertValidatorToRow(validator *types.Validator) *validatorDbRow {
	return &validatorDbRow{
		ConsensusAddress: validator.ConsAddr,
		ConsensusPubKey:  validator.ConsPubKey,
	}
}

// HasValidator implements database.Database
func (db *Database) HasValidator(addr string) (bool, error) {
	var res bool
	stmt := `SELECT EXISTS(SELECT 1 FROM validators WHERE consensus_address = $1);`
	err := db.SQL.QueryRow(stmt, addr).Scan(&res)
	return res, err
}

// SaveValidators implements database.Database
func (db *Database) SaveValidators(validators []*types.Validator) error {
	rows := utils.Map(validators, convertValidatorToRow)

	if len(rows) == 0 {
		return nil
	}

	stmt := `
INSERT INTO validators (consensus_address, consensus_pubkey) 
VALUES (:consensus_address, :consensus_pubkey)
ON CONFLICT DO NOTHING`
	_, err := db.SQL.NamedExec(stmt, rows)
	return err
}
