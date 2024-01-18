package sqlsettings

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/golangast/switchterm/loggers"

	"github.com/golangast/switchterm/db/sqlite/dbconn"
)

func UpdateSettings(dir string) {
	logger := loggers.CreateLogger()

	//opening database
	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"trying to connect to database",
			slog.String("error: ", err.Error()),
		)
	}

	//prepare statement so that no sql injection
	stmt, err := db.Prepare("update settings set dir=?")
	if err != nil {
		logger.Error(
			"trying to prepare update tag in db",
			slog.String("error: ", err.Error()),
		)
	}

	//execute qeury
	_, err = stmt.Exec(dir)
	if err != nil {
		logger.Error(
			"trying to execute db statement",
			slog.String("error: ", err.Error()),
		)
	}

}

func GetDir() (string, error) {
	var id int
	var dir string
	db, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
		return dir, err
	}
	rows, err := db.Query("SELECT * FROM settings")
	if err != nil {
		fmt.Println(err)
		return dir, err
	}
	for rows.Next() {
		err := rows.Scan(&id, &dir)
		if err != nil {
			fmt.Println(err)
			return dir, err
		}

	}
	//close everything
	defer rows.Close()
	defer db.Close()
	switch err {
	case sql.ErrNoRows:
		return dir, err
	case nil:
		return dir, err
	default:
		return dir, err
	}
}
