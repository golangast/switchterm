package domain

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/db/dbconn"
)

func (u *Domains) Create() error {
	logger := loggers.CreateLogger()

	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"conning db",
			slog.String("error: ", err.Error()),
		)
	}

	// Create a statement to insert data into the `users` table.
	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO `domains` (`domain`, `github`) VALUES (?, ?)")
	if err != nil {
		logger.Error(
			"insert domain",
			slog.String("error: ", err.Error()),
		)
	}
	defer stmt.Close()

	// Insert data into the `users` table.
	_, err = stmt.ExecContext(context.Background(), u.Domain, u.Github)
	if err != nil {
		logger.Error(
			"executing insert domain",
			slog.String("error: ", err.Error()),
		)
	}

	db.Close()
	return nil
}
func (u *Domains) GetDomain() (Domains, error) {
	logger := loggers.CreateLogger()

	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"connecting db for TLS",
			slog.String("error: ", err.Error()),
		)
	}
	var (
		id     int
		domain string
		github string
	)
	// get from database
	stmt, err := db.Prepare("SELECT * FROM domains WHERE id = 1")
	if err != nil {
		logger.Error(
			"Select statement db for TLS",
			slog.String("error: ", err.Error()),
		)
	}
	err = stmt.QueryRow(&id).Scan(&id, &domain, &github)
	if err != nil {
		logger.Error(
			"querying db for TLS",
			slog.String("error: ", err.Error()),
		)
	}
	d := Domains{Domain: domain, Github: github}

	defer db.Close()
	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		logger.Error(
			"no rows db for TLS",
			slog.String("error: ", sql.ErrNoRows.Error()),
		)
		return d, nil

	case nil:
		logger.Error(
			"nil rows db for TLS",
			slog.String("error: ", sql.ErrNoRows.Error()),
		)
		return d, nil

	default:
		return d, nil
	}
}

type Domains struct {
	Domain string
	Github string
}