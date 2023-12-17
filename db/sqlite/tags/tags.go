package tags

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/golangast/switchterm/db/sqlite/dbconn"
)

func GetNoteByChosen(cmds []string) ([]Tags, error) {
	var tags []Tags
	db, err := dbconn.DbConnection()
	if err != nil {
		return tags, err
	}

	// get from database
	stmt, err := db.Prepare("SELECT * FROM tags WHERE cmd = ?")
	if err != nil {
		return tags, err
	}

	var id, cmd, note, tag string

	for _, v := range cmds {
		err = stmt.QueryRow(v).Scan(&id, &cmd, &note, &tag)
		if err != nil {
			return tags, err
		}
		t := Tags{ID: id, CMD: cmd, Note: note, Tag: tag}
		tags = append(tags, t)
	}

	defer db.Close()
	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		// close db when not in use
		return tags, nil

	case nil:
		fmt.Println("nil!!!!!!!!!!!!")

		// close db when not in use
		return tags, nil

	default:

		fmt.Println("default!!!!!!!!!!!!")

		return tags, nil
	}
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

func GetCMD() (Tags, error) {
	db, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)

	}
	var (
		id   string
		cmd  string
		note string
		tag  string

		t Tags
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

func GetAll() ([]Tags, error) {
	var (
		id   string
		cmd  string
		note string
		tag  string

		ts []Tags
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
func GetCMDByTag(tag string) ([]string, error) {
	var (
		id   int
		cmd  string
		note string

		tt []string
	)
	db, err := dbconn.DbConnection()
	if err != nil {
		return tt, err
	}

	rows, err := db.Query("SELECT * FROM tags WHERE tag = ?", tag)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return tt, err
	}
	defer rows.Close()
	defer db.Close()

	for rows.Next() {

		if err := rows.Scan(&id, &cmd, &note, &tag); err != nil {
			return tt, err
		}
		fmt.Println(id, cmd, note, tag)
		tt = append(tt, cmd)
	}
	if err := rows.Err(); err != nil {
		return tt, err
	}
	return tt, nil

}

func DeleteTag(cmd string) error {

	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}
	res, err := db.Exec("DELETE FROM tags WHERE cmd =$1", cmd)
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

	return nil
}

type Tags struct {
	ID   string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	CMD  string `valid:"type(string),required" param:"cmd" query:"cmd" form:"cmd" json:"cmd" xml:"cmd" validate:"required,cmd" mod:"trim"`
	Note string `valid:"type(string),required" param:"note" query:"note" form:"note" json:"note" xml:"note"`
	Tag  string `valid:"type(string)" param:"tag" query:"tag" form:"tag" json:"tag" xml:"tag" validate:"required" mod:"trim"`
}
