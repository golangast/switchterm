package createtable

import (
	"fmt"
	"os"

	"github.com/golangast/endrulats/internal/dbsql/dbconn"
)

func CreateTable() error {
	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}
	_, err = db.Query("CREATE TABLE `comment` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `email` VARCHAR(64) NULL, `language` VARCHAR(255) NOT NULL, `comment` VARCHAR(255) NOT NULL, `sitetoken` VARCHAR(255) NULL )")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = db.Query("CREATE TABLE `users` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `email` VARCHAR(64) NULL, `passwordhash` VARCHAR(255) NOT NULL, `passwordraw` VARCHAR(255) NOT NULL, `isdisabled` VARCHAR(255) NULL, `sessionkey` VARCHAR(255) NULL, `sessionname` VARCHAR(255) NULL, `sessiontoken` VARCHAR(255) NULL, `sitetoken` VARCHAR(255) NULL )")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.Close()

	return nil
}
