package auths

import (
	"agency-banking/util"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries AuthRepository
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../../../..")
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	testQueries = NewAuthRepository(testDB)
	os.Exit(m.Run())
}
