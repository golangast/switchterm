package rundbserver

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func Rundbserver() {

	d := domain.Domains{}

	dd, err := d.GetDomain()
	switchutility.Checklogger(err, "trying to get domains for running database server")

	var do []string
	for _, v := range dd {
		do = append(do, v.Domain)
	}

	chosendomain := switchselector.MenuInstuctions(do, 1, "purple", "purple", "Which website are you going to run the database server for?")
	rr := Rander()
	rrr := Rander()

	if err := switchutility.ReplaceLine(chosendomain+"/assets/db/rqlite/bin/config.json", `username`, `"username": "`+rr+`",`); err != nil {
		switchutility.Checklogger(err, "trying to update config.json")
	}
	if err := switchutility.ReplaceLine(chosendomain+"/assets/db/rqlite/bin/config.json", `password`, `"password": "`+rrr+`",`); err != nil {
		switchutility.Checklogger(err, "trying to update config.json")
	}
	if err := switchutility.ReplaceLine(chosendomain+"/internal/dbsql/dbconn/dbconn.go", `conn, err := gorqlite.Open(`, `conn, err := gorqlite.Open("http://`+rr+`:`+rrr+`@localhost:4001/")`); err != nil {
		switchutility.Checklogger(err, "trying to update config.json")
	}

	fmt.Println("Database server is running....open a new terminal if you want to run the server or hit ctrl+c to cancel the db server (this will only work locally!)")
	colortermer.ColorizeOutPut("dpurple", "purple", "the process is called rqlited in your resource manager\n")
	colortermer.ColorizeOutPut("dpurple", "purple", "if you want resources on the db server lookup the following\n")
	colortermer.ColorizeOutPut("dpurple", "purple", "https://github.com/rqlite/rqlite and https://github.com/rqlite/gorqlite\n")

	if err := switchutility.ShellBash("cd "+chosendomain+"/assets/db/rqlite/bin && chmod 755 ./rqlited && ./rqlited -auth config.json  ~/node.1 ", "trying to run database server bash command"); err != nil {
		switchutility.Checklogger(err, "running database server")
	}

}
func Rander() string {
	randomNumber := rand.Intn(1000)
	Randnum := strconv.Itoa(randomNumber)

	return fmt.Sprintf(Randnum)
}
