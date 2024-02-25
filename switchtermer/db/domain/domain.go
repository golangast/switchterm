package domain

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/db/dbconn"
)

func (u *Domains) Create() error {

	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}

	// Create a statement to insert data into the `users` table.
	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO `domains` (`domain`, `github`) VALUES (?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	// Insert data into the `users` table.
	_, err = stmt.ExecContext(context.Background(), u.Domain, u.Github)
	if err != nil {
		return err
	}

	db.Close()
	return nil
}
func (u *Domains) GetDomain() (Domains, error) {
	logger := loggers.CreateLogger()

	db, err := dbconn.DbConnection()
	if err != nil {
		return *u, err
	}

	var (
		id     int
		domain string
		github string
	)
	// get from database
	stmt, err := db.Prepare("SELECT * FROM domains")
	if err != nil {
		return *u, err
	}
	err = stmt.QueryRow(&id).Scan(&id, &domain, &github)
	if err != nil {
		return *u, err
	}
	d := Domains{Domain: domain, Github: github}

	defer db.Close()
	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		logger.Error(
			"no rows db for getting domain",
			slog.String("error: ", sql.ErrNoRows.Error()),
		)
		return d, nil

	case nil:
		logger.Error(
			"nil rows db for getting domain",
			slog.String("error: ", sql.ErrNoRows.Error()),
		)
		return d, nil

	default:
		return d, nil
	}
}

func (u *Domains) GetGitByDomain() (string, error) {
	logger := loggers.CreateLogger()

	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"connecting to db for domains",
			slog.String("error: ", err.Error()),
		)
	}
	var (
		id     int
		domain string
		github string
	)
	// get from database
	stmt, err := db.Prepare("SELECT * FROM domains WHERE domain = ?")
	if err != nil {
		logger.Error(
			"Select statement db for getting githbu",
			slog.String("error: ", err.Error()),
		)
	}
	err = stmt.QueryRow(&u.Domain).Scan(&id, &domain, &github)
	if err != nil {
		logger.Error(
			"querying db for github",
			slog.String("error: ", err.Error()),
		)
	}

	defer db.Close()
	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		logger.Error(
			"no rows db for getting github",
			slog.String("error: ", sql.ErrNoRows.Error()),
		)
		return github, nil

	case nil:

		return github, nil

	default:
		return github, nil
	}
}

type Domains struct {
	Domain string
	Github string
}

func (u *Domains) GetAllDomain() ([]Domains, error) {
	logger := loggers.CreateLogger()

	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"connecting to db for domains",
			slog.String("error: ", err.Error()),
		)
	}
	var (
		id      int
		domain  string
		github  string
		domains []Domains
	)
	rows, err := db.Query("SELECT * FROM domains")
	if err != nil {
		logger.Error(
			"Select statement db for domains rows",
			slog.String("error: ", err.Error()),
		)
	}
	for rows.Next() {

		err := rows.Scan(&id, &domain, &github)
		if err != nil {
			logger.Error(
				"scanning rows for domains",
				slog.String("error: ", err.Error()),
			)
		}

		d := Domains{Domain: domain, Github: github}
		domains = append(domains, d)

	}

	defer db.Close()
	defer rows.Close()
	switch err {
	case sql.ErrNoRows:
		logger.Error(
			"no rows db for domains",
			slog.String("error: ", sql.ErrNoRows.Error()),
		)
		return domains, nil

	case nil:
		logger.Error(
			"nil rows db for domains",
			slog.String("error: ", sql.ErrNoRows.Error()),
		)
		return domains, nil

	default:
		return domains, nil
	}
}
