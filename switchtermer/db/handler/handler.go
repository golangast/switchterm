package handler

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/db/dbconn"
)

type Handler struct {
	Domain  string
	Handle  string
	Segment string
}

func (u *Handler) Create() error {
	logger := loggers.CreateLogger()

	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"conning db for handler",
			slog.String("error: ", err.Error()),
		)
	}

	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO `handler` (`domain`, `handle`, `segment`) VALUES (?, ?, ?)")
	if err != nil {
		logger.Error(
			"create handler",
			slog.String("error: ", err.Error()),
		)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(context.Background(), u.Domain, u.Handle, u.Segment)
	if err != nil {
		logger.Error(
			"executing insert handler",
			slog.String("error: ", err.Error()),
		)
	}

	db.Close()
	return nil
}

func (u *Handler) GetAll() ([]Handler, error) {
	logger := loggers.CreateLogger()

	var id, domain, handle, segment string
	var wholehandler []Handler
	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"conning db for getting all data",
			slog.String("error: ", err.Error()),
		)
	}

	rows, err := db.Query("SELECT * FROM handler")
	if err != nil {
		logger.Error(
			"select all data",
			slog.String("error: ", err.Error()),
		)
	}
	for rows.Next() {
		err := rows.Scan(&id, &domain, &handle, &segment)
		if err != nil {
			logger.Error(
				"going through rows",
				slog.String("error: ", err.Error()),
			)
		}
		t := Handler{Domain: domain, Handle: handle, Segment: segment}
		wholehandler = append(wholehandler, t)
	}
	//close everything
	defer rows.Close()
	defer db.Close()
	switch err {
	case sql.ErrNoRows:
		return wholehandler, err
	case nil:
		return wholehandler, nil
	default:
		return wholehandler, nil
	}
}

func GetStringHandlers() ([]string, error) {
	h := Handler{}
	var hs []string

	hh, err := h.GetAll()
	if err != nil {
		return hs, err
	}

	for _, v := range hh {
		hs = append(hs, v.Handle)
	}

	return hs, nil
}
