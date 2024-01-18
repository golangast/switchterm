package window

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	"github.com/golangast/switchterm/loggers"

	"github.com/golangast/switchterm/db/sqlite/dbconn"
)

func GetTagByName(window []string) ([]Window, error) {
	var windows []Window
	db, err := dbconn.DbConnection()
	if err != nil {
		return windows, err
	}
	stmt, err := db.Prepare("SELECT * FROM window WHERE name = ?")
	if err != nil {
		return windows, err
	}
	var id, name, tag string
	for _, v := range name {
		err = stmt.QueryRow(v).Scan(&id, &name, &tag)
		if err != nil {
			return windows, err
		}
		t := Window{ID: id, Name: name, Tag: tag}

		windows = append(windows, t)
	}
	defer db.Close()
	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		return windows, nil
	case nil:
		return windows, nil
	default:
		return windows, nil
	}
}
func Create(name, tag string) error {
	var err error
	logger := loggers.CreateLogger()
	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"trying to open a database connection",
			slog.String("error: ", err.Error()),
		)
	}
	stmt, err := db.Prepare("INSERT INTO `Window` ( `name`, `tag`) VALUES ( ?,?)")
	if err != nil {
		logger.Error(
			"trying to prepare db statement",
			slog.String("error: ", err.Error()),
		)
	}

	_, err = stmt.Exec(name, tag)
	if err != nil {
		logger.Error(
			"trying to execute db statement",
			slog.String("error: ", err.Error()),
		)
	}
	defer stmt.Close()
	defer db.Close()
	return nil
}

func GetWindowByName(name string) (Window, error) {
	db, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
	}
	var id, tag string
	var t Window
	stmt, err := db.Prepare("SELECT * FROM window WHERE name = ?")
	if err != nil {
		return t, err
	}

	err = stmt.QueryRow(name).Scan(&id, &name, &tag)
	if err != nil {
		return t, err
	}
	tt := Window{ID: id, Name: name, Tag: tag}

	defer db.Close()
	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		return tt, nil
	case nil:
		return tt, nil
	default:
		return tt, nil
	}
}

func GetAll() ([]Window, error) {
	var id, name, tag string
	var ts []Window
	db, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
		return ts, err
	}
	rows, err := db.Query("SELECT * FROM Window")
	if err != nil {
		fmt.Println(err)
		return ts, err
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &tag)
		if err != nil {
			fmt.Println(err)
			return ts, err
		}
		t := Window{ID: id, Name: name, Tag: tag}
		ts = append(ts, t)
	}
	//close everything
	defer rows.Close()
	defer db.Close()
	switch err {
	case sql.ErrNoRows:
		return ts, nil
	case nil:
		return ts, nil
	default:
		return ts, nil
	}
}

// https://golangbot.com/mysql-select-single-multiple-rows/
func GetNameByTag(tag string) ([]string, error) {
	var id, name string
	var tt []string
	db, err := dbconn.DbConnection()
	if err != nil {
		return tt, err
	}
	rows, err := db.Query("SELECT * FROM Window WHERE tag = ?", tag)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return tt, err
	}
	defer rows.Close()
	defer db.Close()
	for rows.Next() {
		if err := rows.Scan(&id, &name, &tag); err != nil {
			return tt, err
		}
		tt = append(tt, name)
	}
	if err := rows.Err(); err != nil {
		return tt, err
	}
	return tt, nil
}

func DeleteTag(name string) error {
	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}
	res, err := db.Exec("DELETE FROM Window WHERE name =$1", name)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		fmt.Println("rows affected were 0!!")
	}
	defer db.Close()

	return nil
}

func UpdateTag(name, tag string) {
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
	stmt, err := db.Prepare("update window set tag=? where name=?")
	if err != nil {
		logger.Error(
			"trying to prepare update window in db",
			slog.String("error: ", err.Error()),
		)
	}

	//execute qeury
	_, err = stmt.Exec(tag, name)
	if err != nil {
		logger.Error(
			"trying to execute db statement",
			slog.String("error: ", err.Error()),
		)
	}

}

type Window struct {
	ID   string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	Name string `valid:"type(string),required" param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required,name" mod:"trim"`
	Tag  string `valid:"type(string),required" param:"tag" query:"tag" form:"tag" json:"tag" xml:"tag"`
}
