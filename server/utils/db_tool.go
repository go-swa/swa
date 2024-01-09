package utils

import (
	"database/sql"
	"github.com/pkg/errors"
)


func DbPing(dbType string, dsn string) error {
	var conn *sql.DB
	var err error

	if dbType == "" || dsn == "" {
		return errors.Errorf("dbType为空 或 dsn为空")
	}

	conn, err = sql.Open(dbType, dsn)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	defer conn.Close()

	return nil
}

