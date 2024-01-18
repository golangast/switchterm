package data

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/golangast/switchterm/db/sqlite/dbconn"
	"github.com/golangast/switchterm/loggers"
)

func GetDataByName(name string) (Data, error) {
	var data Data
	var id, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10 string

	db, err := dbconn.DbConnection()
	if err != nil {
		return data, err
	}
	stmt, err := db.Prepare("SELECT * FROM data WHERE name = ?")
	if err != nil {
		return data, err
	}
	err = stmt.QueryRow(name).Scan(&id, &name, &f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9, &f10)
	if err != nil {
		return data, err
	}
	t := Data{ID: id, Name: name, F1: f1, F2: f2, F3: f3, F4: f4, F5: f5, F6: f6, F7: f7, F8: f8, F9: f9, F10: f10}

	defer db.Close()
	defer stmt.Close()
	switch err {
	case nil:
		return t, nil
	default:
		return t, nil
	}
}
func Create(fields []any) error {
	fmt.Println(len(fields))
	var err error
	logger := loggers.CreateLogger()
	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"trying to open a database connection",
			slog.String("error: ", err.Error()),
		)
	}
	stmt, err := db.Prepare("INSERT INTO `data` ( `id`,`name`,`f1`,`f2`,`f3`,`f4`,`f5`,`f6`,`f7`,`f8`,`f9`,`f10`) VALUES ( NULL, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		logger.Error(
			"trying to prepare db statement",
			slog.String("error: ", err.Error()),
		)
	}

	_, err = stmt.Exec(fields...)
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
func GetAll() ([]Data, error) {
	var id, name, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10 string
	var da []Data
	db, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
		return da, err
	}
	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		fmt.Println(err)
		return da, err
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9, &f10)
		if err != nil {
			fmt.Println(err)
			return da, err
		}
		t := Data{ID: id, Name: name, F1: f1, F2: f2, F3: f3, F4: f4, F5: f5, F6: f6, F7: f7, F8: f8, F9: f9, F10: f10}
		da = append(da, t)
	}
	//close everything
	defer rows.Close()
	defer db.Close()
	switch err {
	case sql.ErrNoRows:
		return da, nil
	case nil:
		return da, nil
	default:
		return da, nil
	}
}
func DeleteData(name string) error {
	db, err := dbconn.DbConnection()
	if err != nil {
		return err
	}
	res, err := db.Exec("DELETE FROM data WHERE name =$1", name)
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

func UpdateData(name string, fields []string) {
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
	stmt, err := db.Prepare("update data set tag=? where tag=?")
	if err != nil {
		logger.Error(
			"trying to prepare update tag in db",
			slog.String("error: ", err.Error()),
		)
	}

	//execute qeury
	_, err = stmt.Exec(name)
	if err != nil {
		logger.Error(
			"trying to execute db statement",
			slog.String("error: ", err.Error()),
		)
	}

}

type TableData struct {
	Name    []string
	Columns []string
	Values  []string
}

func Getapptables() []TableData {
	logger := loggers.CreateLogger()

	db, err := dbconn.DbConnection()
	if err != nil {
		logger.Error(
			"trying to connect to database",
			slog.String("error: ", err.Error()),
		)
	}

	var (
		tables []string
		types  string
		name   string
		TDS    []TableData
	)

	rows, err := db.Query("SELECT type, name FROM sqlite_master where type='table'")
	if err != nil {
		logger.Error(
			"trying to qeury sqlite_master",
			slog.String("error: ", err.Error()),
		)
	}

	//cycle through the rows to collect all the data
	for rows.Next() {

		err := rows.Scan(&types, &name)
		if err != nil {
			logger.Error(
				"trying to scan rows",
				slog.String("error: ", err.Error()),
			)
		}

		//store into memory
		tables = append(tables, name)

	}
	//after table names have been appended
	//grab their data.
	for _, table := range tables {
		TD := TableData{}
		rows, err := db.Query("SELECT * FROM " + table + ";")
		if err != nil {
			logger.Error(
				"trying to select table",
				slog.String("error: ", err.Error()),
			)
		}
		columns, err := rows.Columns()
		if err != nil {
			logger.Error(
				"trying to scan coluns",
				slog.String("error: ", err.Error()),
			)
		}
		count := len(columns)
		values := make([]any, count)
		valuePtr := make([]any, count)
		var v any
		var prevtable string
		for rows.Next() {

			if prevtable != table {
				//fmt.Print(" table:", table, " columns:", columns)
				//scan needs any type so turn columns into []any
				for i, _ := range columns {
					valuePtr[i] = &values[i]
				}
				prevtable = table
				err := rows.Scan(valuePtr...)
				if err != nil {
					logger.Error(
						"trying to scan table values",
						slog.String("error: ", err.Error()),
					)
				}
				//go through the columns
				for a, _ := range columns {

					val := values[a]

					b, ok := val.([]byte)

					if ok {
						v = string(b)
					} else {
						v = val
					}

					TD.Values = append(TD.Values, fmt.Sprint(v))
				}
			}

		}
		TD.Columns = columns
		TD.Name = append(TD.Name, table)
		TDS = append(TDS, TD)
	}

	//close everything
	defer rows.Close()
	defer db.Close()
	return TDS

}

type Data struct {
	ID   string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	Name string `valid:"type(string),required" param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required,name" mod:"trim"`
	F1   string `valid:"type(string)" param:"f1" query:"f1" form:"f1" json:"f1" xml:"f1"`
	F2   string `valid:"type(string)" param:"f2" query:"f2" form:"f2" json:"f2" xml:"f2"`
	F3   string `valid:"type(string)" param:"f3" query:"f3" form:"f3" json:"f3" xml:"f3"`
	F4   string `valid:"type(string)" param:"f4" query:"f4" form:"f4" json:"f4" xml:"f4"`
	F5   string `valid:"type(string)" param:"f5" query:"f5" form:"f5" json:"f5" xml:"f5"`
	F6   string `valid:"type(string)" param:"f6" query:"f6" form:"f6" json:"f6" xml:"f6"`
	F7   string `valid:"type(string)" param:"f7" query:"f7" form:"f7" json:"f7" xml:"f7"`
	F8   string `valid:"type(string)" param:"f8" query:"f8" form:"f8" json:"f8" xml:"f8"`
	F9   string `valid:"type(string)" param:"f9" query:"f9" form:"f9" json:"f9" xml:"f9"`
	F10  string `valid:"type(string)" param:"f10" query:"f10" form:"f10" json:"f10" xml:"f10"`
}
