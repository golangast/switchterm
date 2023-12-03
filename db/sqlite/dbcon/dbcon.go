package dbcon

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

const file string = "./data.db"

func DbConnection() (*sql.DB, error) {
	fmt.Println(file)

	db, err := sql.Open("sqlite", file)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
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
