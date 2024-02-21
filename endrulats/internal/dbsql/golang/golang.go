
	package golang
	
	import (
		"fmt"
		"github.com/golangast/endrulats/internal/dbsql/dbconn"
	)
	
	
	type Golang struct {
		ID     int   `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id" `
		
		Age int  `param:"Age" query:"Age" header:"Age" form:"Age" json:"Age" xml:"Age" `
		
		Email string  `param:"Email" query:"Email" header:"Email" form:"Email" json:"Email" xml:"Email" `
		
		Name string  `param:"Name" query:"Name" header:"Name" form:"Name" json:"Name" xml:"Name" `
		
		
		
		}
	
	func GetGolang() []Golang {

	data, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
	}

	//variables used to store data from the query
	var (
		id int
		
		age int
		
		email string
		
		name string
		
		golangs  []Golang 
		)//used to store all users
		//https://go.dev/play/p/82imTtvHWzb
	_, err = data.Query("CREATE TABLE IF NOT EXISTS golang (id INTEGER PRIMARY KEY AUTOINCREMENT,  age int NULL  , email text NULL  , name text NULL )")
	if err != nil {
		fmt.Println(err)
	}
	
	//get from database
	rows, err := data.Query("select * from golang")
	if err != nil {
		fmt.Println(err)
	}

	
	
	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &name,&age,&email )
		if err != nil {
			fmt.Println(err)
		}
		
		u := Golang{ID: id,
			
			Age: age,
			
			Email: email,
			
			Name: name,
			
			} 
			golangs = append(golangs, u)
		}
	
	
	//close everything
	rows.Close()
	data.Close()
	return golangs
	}
	