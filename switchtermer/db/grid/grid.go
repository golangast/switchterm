package grid

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/db/dbconn"
)

type Grid struct {
	Name       string
	Domain     string
	Handler    string
	GridLayout string
}

func (t *Grid) Exists(Grid string) (bool, error) {
	var exists bool
	db, err := dbconn.DbConnection()
	if err != nil {
		return false, err
	}
	stmts := db.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM grid WHERE name=?)", t.Name)
	err = stmts.Scan(&exists)
	if err != nil {
		return false, err
	}
	db.Close()

	return exists, nil

}
func (u *Grid) Create() error {

	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}

	// Create a statement to insert data into the `users` table.
	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO `grid` (`name`, `domain`, `handler`, `gridLayout`) VALUES (?, ?,?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	// Insert data into the `grid` table.
	_, err = stmt.ExecContext(context.Background(), u.Name, u.Domain, u.Handler, u.GridLayout)
	if err != nil {
		return err
	}

	db.Close()
	return nil
}
func (u *Grid) GetGrid() ([]Grid, error) {
	logger := loggers.CreateLogger()
	var grids []Grid

	db, err := dbconn.DbConnection()
	if err != nil {
		return grids, err
	}

	var id int

	rows, err := db.Query("SELECT * FROM grid")
	if err != nil {
		logger.Error(
			"Select statement db for domains rows",
			slog.String("error: ", err.Error()),
		)
	}

	for rows.Next() {

		err := rows.Scan(&id, &u.Name, &u.Domain, &u.Handler, &u.GridLayout)
		if err != nil {
			logger.Error(
				"scanning rows for domains",
				slog.String("error: ", err.Error()),
			)
		}
		d := Grid{Name: u.Name, Domain: u.Domain, Handler: u.Handler, GridLayout: u.GridLayout}

		grids = append(grids, d)

	}

	defer db.Close()
	defer rows.Close()
	switch err {
	case sql.ErrNoRows:
		logger.Error(
			"no rows db for getting domain",
			slog.String("error: ", sql.ErrNoRows.Error()),
		)
		return grids, nil

	case nil:
		logger.Error(
			"nil rows db for getting domain",
			slog.String("error: ", sql.ErrNoRows.Error()),
		)
		return grids, nil

	default:
		return grids, nil
	}
}

func (u Grid) GetGridByDomain() (Grid, error) {
	logger := loggers.CreateLogger()

	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"connecting to db for domains",
			slog.String("error: ", err.Error()),
		)
	}
	var id int
	// get from database
	stmt, err := db.Prepare("SELECT * FROM grid WHERE domain = ?")
	if err != nil {
		logger.Error(
			"Select statement db for getting githbu",
			slog.String("error: ", err.Error()),
		)
	}
	err = stmt.QueryRow(&u.Domain).Scan(&id, &u.Name, &u.Domain, &u.Handler, &u.GridLayout)
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
		return u, nil

	case nil:

		return u, nil

	default:
		return u, nil
	}
}

func (u *Grid) GetAllGrids() ([]Grid, error) {
	logger := loggers.CreateLogger()

	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"connecting to db for domains",
			slog.String("error: ", err.Error()),
		)
	}
	var (
		id    int
		grids []Grid
	)
	rows, err := db.Query("SELECT * FROM grid")
	if err != nil {
		logger.Error(
			"Select statement db for domains rows",
			slog.String("error: ", err.Error()),
		)
	}
	for rows.Next() {
		err = rows.Scan(&id, &u.Name, &u.Domain, &u.Handler, &u.GridLayout)

		if err != nil {
			logger.Error(
				"scanning rows for domains",
				slog.String("error: ", err.Error()),
			)
		}

		d := Grid{Name: u.Name, Domain: u.Domain, Handler: u.Handler, GridLayout: u.GridLayout}
		grids = append(grids, d)

	}

	defer db.Close()
	defer rows.Close()
	switch err {
	case sql.ErrNoRows:
		return grids, sql.ErrNoRows
	case nil:
		return grids, nil

	default:
		return grids, nil
	}
}

func GetStringGrids() ([]string, error) {
	d := Grid{}
	var do []string

	dd, err := d.GetGrid()
	if err != nil {
		return do, err
	}

	for _, v := range dd {
		do = append(do, v.Name)
	}

	return do, nil
}
