package gar

import (
	"fmt"

	"github.com/golangast/endrulats/internal/dbsql/dbconn"
)

type Gar struct {
	ID int `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id" `

	Age int `param:"Age" query:"Age" header:"Age" form:"Age" json:"Age" xml:"Age" `

	Name string `param:"Name" query:"Name" header:"Name" form:"Name" json:"Name" xml:"Name" `
}

func GetGar() []Gar {

	data, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
	}

	//variables used to store data from the query
	var (
		id int

		age int

		name string

		gars []Gar
	) //used to store all users
	_, err = data.Query("CREATE TABLE IF NOT EXISTS gar (id INTEGER PRIMARY KEY AUTOINCREMENT,  age int NULL  , name text NULL )")
	if err != nil {
		fmt.Println(err)
	}

	//get from database
	rows, err := data.Query("select * from gar")
	if err != nil {
		fmt.Println(err)
	}

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println(err)
		}

		u := Gar{ID: id,

			Age: age,

			Name: name,
		}
		gars = append(gars, u)
	}

	//close everything
	rows.Close()
	data.Close()
	return gars
}
