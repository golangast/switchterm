package createtable

import (
	"fmt"
	"os"

	"github.com/golangast/switchterm/db/sqlite/dbconn"
)

func CreateTable() error {
	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}
	_, err = db.Query("CREATE TABLE `domains` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `domain` VARCHAR(64) NULL, `github` VARCHAR(255) NOT NULL)")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.Close()

	return nil
}
