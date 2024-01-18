package tags

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	"github.com/golangast/switchterm/loggers"

	"github.com/golangast/switchterm/db/sqlite/dbconn"
)

func GetNoteByChosen(cmds []string) ([]Tags, error) {

	var tags []Tags
	var id, cmd, note, tag, bash, bashfile, bashdir string

	db, err := dbconn.DbConnection()
	if err != nil {
		return tags, err
	}
	stmt, err := db.Prepare("SELECT * FROM tags WHERE cmd = ?")
	if err != nil {
		return tags, err
	}
	for _, v := range cmds {
		err = stmt.QueryRow(v).Scan(&id, &cmd, &note, &tag, &bash, &bashfile, &bashdir)
		if err != nil {
			return tags, err
		}
		t := Tags{ID: id, CMD: cmd, Note: note, Tag: tag, Bash: bash, Bashfile: bashfile, Bashdir: bashdir}
		tags = append(tags, t)
	}
	defer db.Close()
	defer stmt.Close()
	switch err {

	case nil:
		return tags, nil
	default:
		return tags, nil
	}
}
func Create(cmd, note, tag, bash, bashdir, bashfile string) error {

	var err error
	logger := loggers.CreateLogger()
	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"trying to open a database connection",
			slog.String("error: ", err.Error()),
		)
	}
	stmt, err := db.Prepare("INSERT INTO `tags` ( `cmd`, `note`, `tag`, `bash`, `bashfile`, `bashdir`) VALUES ( ?,?, ?,?, ?, ?)")
	if err != nil {
		logger.Error(
			"trying to prepare db statement",
			slog.String("error: ", err.Error()),
		)
	}

	_, err = stmt.Exec(cmd, note, tag, bash, bashfile, bashdir)
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

func GetCMD() (Tags, error) {
	db, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
	}
	var id, cmd, note, tag, bash, bashfile, bashdir string
	var t Tags
	stmt, err := db.Prepare("SELECT * FROM tags WHERE cmd = ?")
	if err != nil {
		return t, err
	}

	err = stmt.QueryRow(cmd).Scan(&id, &cmd, &note, &tag, &bash, &bashfile, &bashdir)
	if err != nil {
		return t, err
	}
	t = Tags{ID: id, CMD: cmd, Note: note, Tag: tag, Bash: bash, Bashfile: bashfile, Bashdir: bashdir}

	defer db.Close()
	defer stmt.Close()
	switch err {
	case sql.ErrNoRows:
		return t, nil
	case nil:
		return t, nil
	default:
		return t, nil
	}
}

func GetAll() ([]Tags, error) {
	var id, cmd, note, tag, bash, bashfile, bashdir string
	var ts []Tags
	db, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
		return ts, err
	}
	rows, err := db.Query("SELECT * FROM tags")
	if err != nil {
		fmt.Println(err)
		return ts, err
	}
	for rows.Next() {
		err := rows.Scan(&id, &cmd, &note, &tag, &bash, &bashfile, &bashdir)
		if err != nil {
			fmt.Println(err)
			return ts, err
		}
		t := Tags{ID: id, CMD: cmd, Note: note, Tag: tag, Bash: bash, Bashfile: bashfile, Bashdir: bashdir}
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
func GetCMDByTag(tag string) ([]string, error) {
	var id, cmd, note, bash, bashfile, bashdir string
	var tt []string
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
		if err := rows.Scan(&id, &cmd, &note, &tag, &bash, &bashfile, &bashdir); err != nil {
			return tt, err
		}
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
	stmt, err := db.Prepare("update tags set tag=? where tag=?")
	if err != nil {
		logger.Error(
			"trying to prepare update tag in db",
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

type Tags struct {
	ID       string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	CMD      string `valid:"type(string),required" param:"cmd" query:"cmd" form:"cmd" json:"cmd" xml:"cmd" validate:"required,cmd" mod:"trim"`
	Note     string `valid:"type(string),required" param:"note" query:"note" form:"note" json:"note" xml:"note"`
	Tag      string `valid:"type(string)" param:"tag" query:"tag" form:"tag" json:"tag" xml:"tag" validate:"required" mod:"trim"`
	Bash     string `valid:"type(string)" param:"bash" query:"bash" form:"bash" json:"bash" xml:"bash" validate:"required" mod:"trim"`
	Bashfile string `valid:"type(string)" param:"bashfile" query:"bashfile" form:"bashfile" json:"bashfile" xml:"bashfile" validate:"required" mod:"trim"`
	Bashdir  string `valid:"type(string)" param:"bashdir" query:"bashdir" form:"bashdir" json:"bashdir" xml:"bashdir" validate:"required" mod:"trim"`
}
