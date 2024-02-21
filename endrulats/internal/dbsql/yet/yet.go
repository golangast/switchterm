package yet

import (
	"fmt"
	"os"

	"github.com/golangast/endrulats/internal/dbsql/dbconn"
)

type Yet struct {
	ID int `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id" `

	Age int `param:"Age" query:"Age" header:"Age" form:"Age" json:"Age" xml:"Age" `

	Name string `param:"Name" query:"Name" header:"Name" form:"Name" json:"Name" xml:"Name" `
}

func GetYet() []Yet {

	data, err := dbconn.DbConnection()
	if err != nil {
		fmt.Println(err)
	}

	_, err = data.Query("CREATE TABLE IF NOT EXISTS `yet` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `age` VARCHAR(64) NULL, `name` VARCHAR(255) NOT NULL)")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//variables used to store data from the query
	var (
		id int

		age int

		name string

		yets []Yet //used to store all users
	)

	//get from database
	rows, err := data.Query("select * from yet")
	if err != nil {
		fmt.Println(err)
	}

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println(err)
		}

		u := Yet{ID: id,

			Age: age,

			Name: name,
		}
		yets = append(yets, u)
	}

	//close everything
	rows.Close()
	data.Close()
	return yets
}
