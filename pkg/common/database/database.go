package database

import (
	"github.com/alisyahbana/efishery-test/pkg/common/config"
	"github.com/alisyahbana/efishery-test/pkg/common/env"
	"github.com/alisyahbana/efishery-test/pkg/common/log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	MasterConnString string `json:"masterConnString"`
	SlaveConnString  string `json:"slaveConnString"`
}

var dbMaster *sqlx.DB
var dbSlave *sqlx.DB

func init() {
	var cfg *Config
	config.LoadConfiguration(&cfg, "database", env.GetEnv())
	db, err := sqlx.Open("mysql", cfg.MasterConnString)
	if err != nil {
		panic("Cannot connect to " + cfg.MasterConnString)
	}
	dbMaster = db

	db, err = sqlx.Open("mysql", cfg.SlaveConnString)
	if err != nil {
		panic("Cannot connect to " + cfg.SlaveConnString)
	}
	dbSlave = db
}

func GetDBMaster() *sqlx.DB {
	return dbMaster
}

func GetDBSlave() *sqlx.DB {
	return dbMaster
}

func Prepare(db *sqlx.DB, query string) *sqlx.Stmt {
	stmt, err := db.Preparex(query)
	if err != nil {
		log.Error(err.Error())
	}
	return stmt
}

func StartTransaction(db *sqlx.DB) (*sql.Tx, error) {
	return db.Begin()
}

func CommitTransaction(tx *sql.Tx) error {
	return tx.Commit()
}
