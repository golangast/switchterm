
	package dog
	
	import (
		"fmt"
		"github.com/golangast/endrulats/internal/dbsql/dbconn"
	)
	
	
	type Dog struct {
		ID     int   `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id" `
		
		Name string  `param:"Name" query:"Name" header:"Name" form:"Name" json:"Name" xml:"Name" `
		
		Age int  `param:"age" query:"age" header:"age" form:"age" json:"age" xml:"age" `
		
		
		
		}
	
	func GetDog() []Dog {

	data, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
	}

	//variables used to store data from the query
	var (
		id int
		
		name string
		
		age int
		
		dogs  []Dog //used to store all users
	)
	
	//get from database
	rows, err := data.Query("select * from dog")
	if err != nil {
		fmt.Println(err)
	}
	
	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &name,&age )
		if err != nil {
			fmt.Println(err)
		}
		
		u := Dog{ID: id,
			
			Age: age,
			
			Name: name,
			
			} 
			dogs = append(dogs, u)
		}
	
	
	//close everything
	rows.Close()
	data.Close()
	return dogs
	}
	