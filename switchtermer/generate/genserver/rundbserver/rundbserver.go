package rundbserver

import (
	"fmt"

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
	fmt.Println("Database server is running....open a new terminal if you want to run the server or hit ctrl+c to cancel the db server")
	colortermer.ColorizeOutPut("dpurple", "purple", "the process is called rqlited in your resource manager\n")
	colortermer.ColorizeOutPut("dpurple", "purple", "if you want resources on the db server lookup the following\n")
	colortermer.ColorizeOutPut("dpurple", "purple", "https://github.com/rqlite/rqlite and https://github.com/rqlite/gorqlite\n")

	if err := switchutility.ShellBash("cd "+chosendomain+"/assets/db/rqlite/bin && ./rqlited -auth config.json  ~/node.1 ", "trying to run database server bash command"); err != nil {
		switchutility.Checklogger(err, "running database server")
	}

}
