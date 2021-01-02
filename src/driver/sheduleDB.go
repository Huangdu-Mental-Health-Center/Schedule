package driver

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	ASSESR_DIR = "./assets/"
)

func ConnectDB(nameDB string) (isOK bool, db *sql.DB) {
	db_driver := ASSESR_DIR + nameDB + ".db"
	var isOpen bool
	db, err := sql.Open("sqlite3", db_driver)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	return isOpen, db
}
