package main

import (
	"database/sql"
)

func IsNotFound(err error) bool {
	return err == sql.ErrNoRows
}
