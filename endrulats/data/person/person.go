package person

import (
	"database/sql"
	"fmt"
)

type Person struct {
	ID     int   ` + "`" + `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id" ` + "`" + `
	Name string  ` + "`" + `param:"name" query:"name" header:"name" form:"name" json:"name" xml:"name" ` + "`" + `
	Age     string ` + "`" + `param:"age" query:"age" header:"age" form:"age" json:"age" xml:"age"` + "`" + `
	}

func GetPerson(table string) []Data {
//variables used to store data from the query
var (
	done   string
	dtwo    string
	deleven    string
	Datas  []Data //used to store all users
)

//get from database
rows, err := data.Query("select * from "+table+"")
if err != nil {
	fmt.Println(err)
}

//cycle through the rows to collect all the data
for rows.Next() {
	err := rows.Scan(&id, 
		&done, 
		&dtwo, &dthree, &dfour, &dfive, &dsix, &dseven, &deight, &dnine, &dten, &deleven)
	if err != nil {
		fmt.Println(err)
	}
	//store into memory
	u := DBFields{ID: id, Done: done,  Dtwo: dtwo,  Dthree: dthree,  Dfour: dfour,  Dfive: dfive,  Dsix: dsix,  Dseven: dseven,  Deight: deight,  Dnine: dnine,  Dten: dten,  Deleven: deleven}
	Datas = append(Datas, u)

}
//close everything
rows.Close()
data.Close()
return Datas
}