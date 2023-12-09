package tags

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/golangast/switchterm/db/sqlite/dbconn"
)

func (t *Tags) Exists() error {
	var exists bool
	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}
	stmts := db.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE email=?)", t.CMD)
	err = stmts.Scan(&exists)
	if err != nil {
		return err
	}
	db.Close()

	return nil

}
func Exists(cmd string) (bool, error) {
	db, err := dbconn.DbConnection()
	if err != nil {
		return false, err
	}

	stmts := db.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM tags WHERE cmd=?)", cmd)
	err = stmts.Scan(&cmd)
	if err != nil {
		return false, err
	}

	db.Close()

	return true, nil

}
func Create(cmd, note, tag string) error {

	var id int

	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}
	// Create a statement to insert data into the `tags` table.
	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO `tags` (`id`, `cmd`, `note`, `tag`) VALUES (?, ?,?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Insert data into the `tags` table.
	_, err = stmt.ExecContext(context.Background(), id, cmd, note, tag)
	if err != nil {
		panic(err)
	}

	db.Close()
	return nil
}

func (tags *Tags) GetCMD(cmd string) (Tags, error) {
	db, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
		return *tags, err
	}
	var (
		id   string
		note string
		tag  string
		t    Tags
	)

	//get from database
	stmt, err := db.Prepare("SELECT * FROM tags WHERE cmd = ?")
	if err != nil {
		return t, err
	}
	err = stmt.QueryRow(cmd).Scan(&id, &cmd, &note, &tag)
	if err != nil {
		return t, err
	}
	t = Tags{ID: id, CMD: cmd, Note: note, Tag: tag}
	defer db.Close()
	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		// close db when not in use
		return t, nil

	case nil:
		fmt.Println("nil!!!!!!!!!!!!")

		// close db when not in use
		return t, nil

	default:

		fmt.Println("default!!!!!!!!!!!!")

		return t, nil
	}

}

func (tags *Tags) GetAll() ([]Tags, error) {
	var (
		id   string
		cmd  string
		note string
		tag  string
		ts   []Tags
	)
	db, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
		return ts, err
	}

	//get from database
	rows, err := db.Query("SELECT * FROM tags")
	if err != nil {
		fmt.Println(err)
		return ts, err
	}

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &cmd, &note, &tag)
		if err != nil {
			fmt.Println(err)
			return ts, err
		}

		// //store into memory
		t := Tags{ID: id, CMD: cmd, Note: note, Tag: tag}

		ts = append(ts, t)

	}
	//close everything
	defer rows.Close()
	defer db.Close()
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		// close db when not in use
		return ts, nil

	case nil:
		fmt.Println("nil!!!!!!!!!!!!")
		// close db when not in use
		return ts, nil

	default:
		fmt.Println("default!!!!!!!!!!!!")
		return ts, nil
	}

}

// https://golangbot.com/mysql-select-single-multiple-rows/
func (tags Tags) GetCMDByTag(tag string) (Tags, error) {
	var (
		id   string
		cmd  string
		note string
		t    Tags
	)
	db, err := dbconn.DbConnection()
	if err != nil {
		return t, err
	}

	//get from database
	stmt, err := db.Prepare("SELECT * FROM tags WHERE tag = ?")
	if err != nil {
		return t, err
	}
	err = stmt.QueryRow(tag).Scan(&id, &cmd, &note, &tag)
	if err != nil {
		return t, err
	}
	t = Tags{ID: id, CMD: cmd, Note: note, Tag: tag}
	defer db.Close()
	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return t, nil

	case nil:
		fmt.Println("was nil !!!!!!!!!!!!!", tag)
		return t, nil

	default:
		fmt.Println("default!!!!!!!!!!!!")
		return t, nil
	}

}

type Tags struct {
	ID   string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	CMD  string `valid:"type(string),required" param:"cmd" query:"cmd" form:"cmd" json:"cmd" xml:"cmd" validate:"required,cmd" mod:"trim"`
	Note string `valid:"type(string),required" param:"note" query:"note" form:"note" json:"note" xml:"note"`
	Tag  string `valid:"type(string)" param:"tag" query:"tag" form:"tag" json:"tag" xml:"tag" validate:"required" mod:"trim"`
}
