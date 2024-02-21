package data

import (
	"context"
	"database/sql"
	"log/slog"
	"strings"

	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/db/dbconn"
)

func (u *Fields) Create() error {
	logger := loggers.CreateLogger()

	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"conning db",
			slog.String("error: ", err.Error()),
		)
	}

	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO `data` (`name`, `field`) VALUES (?, ?)")
	if err != nil {
		logger.Error(
			"insert domain",
			slog.String("error: ", err.Error()),
		)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(context.Background(), u.Name, strings.Join(u.Field, " "))
	if err != nil {
		logger.Error(
			"executing insert domain",
			slog.String("error: ", err.Error()),
		)
	}

	db.Close()
	return nil
}

func (u *Fields) GetAll() ([]Fields, error) {
	logger := loggers.CreateLogger()

	var id, name, field string
	var wholefields []Fields
	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"conning db for getting all data",
			slog.String("error: ", err.Error()),
		)
	}

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		logger.Error(
			"select all data",
			slog.String("error: ", err.Error()),
		)
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &field)
		if err != nil {
			logger.Error(
				"going through rows",
				slog.String("error: ", err.Error()),
			)
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
