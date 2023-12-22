package dbconn

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

const file string = "./db/data.db"

func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite", file)
	if err != nil {
		fmt.Println("Error opening", file, err)
		return nil, err
	}
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)
	db.SetConnMaxLifetime(time.Minute * 5)
	//check if it pings
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
} //end of connect
func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
