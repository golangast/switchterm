package main

import (
	"fmt"
	"os"

	"github.com/rqlite/gorqlite"
)

func main() {

	conn, err := gorqlite.Open("http://localhost:4001/") // connects to localhost on 4001 without auth
	if err != nil {
		fmt.Println(err)
	}

	_, err = conn.WriteOne("CREATE TABLE IF NOT EXISTS `test1` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `email` VARCHAR(64) NULL, `language` VARCHAR(255) NOT NULL, `comment` VARCHAR(255) NOT NULL, `sitetoken` VARCHAR(255) NULL )")
	if err != nil {
		fmt.Println(err)
	}
	// CREATE TABLE IF NOT EXISTS TABLE Persons ( PersonID int, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255));
	// using nullable types

	// seq, err := conn.QueueOneParameterized(
	// 	gorqlite.ParameterizedStatement{
	// 		Query:     "INSERT INTO foo(id, name, secret) VALUES(?, ?)",
	// 		Arguments: []interface{}{2, "James Bond"},
	// 	},
	// )
	// qr, err := conn.QueryOneParameterized(
	// 	gorqlite.ParameterizedStatement{
	// 		Query:     "update foo set name=`jimm22` where id=1",
	// 		Arguments: []interface{}{""},
	// 	},
	// )
	// wr, err := conn.WriteParameterized(
	// 	[]gorqlite.ParameterizedStatement{
	// 		{
	// 			Query:     "update foos set name=? where id=1",
	// 			Arguments: []interface{}{"jimm222"},
	// 		},
	// 	},
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }

	rows, err := conn.QueryParameterized(
		[]gorqlite.ParameterizedStatement{
			{
				Query:     "select * from {{.Name}}",
				Arguments: []interface{}{3},
			},
		},
	)
	
	
	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &{{ lower .Fields}} )
		if err != nil {
			fmt.Println(err)
		}
		
		u := {{title .Name}}{ID: id,
			{{range $k, $v := .MapFields }}
			{{$k}}: {{lower $v}},
			{{end}}
			} 
			{{.Name}}s = append({{.Name}}s, u)
		}
	}
	// fmt.Println(qr, err)

	// change my mind and watch the trace
	gorqlite.TraceOn(os.Stderr)
	// fmt.Printf("last insert id was %d\n", res.LastInsertID)
	// stmts := db.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE email=?)", u.Email)
	// err = stmts.Scan(&exists)
	// var id int64
	// var name string
	qr, err := conn.QueryOneParameterized(
		gorqlite.ParameterizedStatement{
			Query:     "SELECT id, name from foo where id > ?",
			Arguments: []interface{}{7},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("query returned %d rows\n", qr.RowNumber())
	for qr.Next() {
		err := qr.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("this is row number %d\n", qr.RowNumber)
		fmt.Printf("there are %d rows overall%d\n", qr.NumRows)
	}

	// for qr.Next() {
	// 	m, err := qr.Map()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	id := m["name"].(float64) // the only json number type
	// 	name := m["name"].(string)

	// 	fmt.Println(id, name)
	// }

}
