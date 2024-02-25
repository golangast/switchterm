package data

import (
	"context"
	"database/sql"
	"strings"

	"github.com/golangast/switchterm/switchtermer/db/dbconn"
)

func (u *Fields) Create() error {
	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}

	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO `data` (`name`, `field`) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(context.Background(), u.Name, strings.Join(u.Field, " "))
	if err != nil {
		return err
	}

	db.Close()
	return nil
}

func (u *Fields) GetAll() ([]Fields, error) {

	var id, name, field string
	var wholefields []Fields
	db, err := dbconn.DbConnection()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &field)
		if err != nil {
			return nil, err
		}
		t := Fields{Name: name, Fields: field}
		wholefields = append(wholefields, t)
	}
	//close everything
	defer rows.Close()
	defer db.Close()
	switch err {
	case sql.ErrNoRows:
		return wholefields, err
	case nil:
		return wholefields, nil
	default:
		return wholefields, nil
	}
}

type Fields struct {
	Name   string
	Fields string
	Field  []string
}
