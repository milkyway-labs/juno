package postgresql_test

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/forbole/juno/v5/database"
	databaseconfig "github.com/forbole/juno/v5/database/config"
	postgres "github.com/forbole/juno/v5/database/postgresql"
	"github.com/forbole/juno/v5/logging"
	"github.com/forbole/juno/v5/types"
)

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}

type DbTestSuite struct {
	suite.Suite

	database *postgres.Database
}

func (suite *DbTestSuite) SetupTest() {
	// Create the codec
	codec := types.MakeTestEncodingConfig()

	// Build the database config
	dbCfg := databaseconfig.DefaultDatabaseConfig().
		WithURL("postgres://user:password@localhost:6433/juno?sslmode=disable&search_path=public")

	// Build the database
	dbCtx := database.NewContext(
		dbCfg,
		codec,
		logging.DefaultLogger(),
		nil,
		types.DefaultTransactionFilter(),
		types.DefaultMessageFilter(),
	)
	db, err := postgres.Builder(dbCtx)
	suite.Require().NoError(err)

	bigDipperDb, ok := (db).(*postgres.Database)
	suite.Require().True(ok)

	// Delete the public schema
	_, err = bigDipperDb.SQL.Exec(`DROP SCHEMA public CASCADE;`)
	suite.Require().NoError(err)

	// Re-create the schema
	_, err = bigDipperDb.SQL.Exec(`CREATE SCHEMA public;`)
	suite.Require().NoError(err)

	dirPath := path.Join(".")
	dir, err := os.ReadDir(dirPath)
	suite.Require().NoError(err)

	for _, fileInfo := range dir {
		if !strings.Contains(fileInfo.Name(), ".sql") {
			continue
		}

		file, err := os.ReadFile(filepath.Join(dirPath, fileInfo.Name()))
		suite.Require().NoError(err)

		commentsRegExp := regexp.MustCompile(`/\*.*\*/`)
		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err := bigDipperDb.SQL.Exec(commentsRegExp.ReplaceAllString(request, ""))
			suite.Require().NoError(err)
		}
	}

	suite.database = bigDipperDb
}
