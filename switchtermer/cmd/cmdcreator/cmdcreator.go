package cmdcreator

import (
	ff "github.com/golangast/gentil/utility/ff"
	temps "github.com/golangast/gentil/utility/temp"

	"github.com/golangast/switchterm/db/sqlite/sqlsettings"
	"github.com/golangast/switchterm/switchtermer/cmd/cmdcreator/temp"
	"github.com/golangast/switchterm/switchtermer/switchutility"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Cmdcreator(cmd string) {

	dir, err := sqlsettings.GetDir()
	switchutility.Checklogger(err, "getting the directory")

	//make file
	cmdfile, err := ff.Filefolder("./cmd/"+cmd, cmd+".go")
	switchutility.Checklogger(err, "trying to create handler file")

	m := make(map[string]string)
	m["cmd"] = cmd
	m["CMD"] = cases.Title(language.Und, cases.NoLower).String(cmd)

	if err := temps.Writetemplate(temp.CmdTemp, cmdfile, m); err != nil {
		switchutility.Checklogger(err, "trying to update router.html")
	}
	if err := switchutility.UpdateText("./switchtermer/cmd/cmdrunner/cmdrunner.go", `case "`+cmd+`"`+":"+"\n"+cmd+"."+cases.Title(language.Und, cases.NoLower).String(cmd)+"()", `//#addcmd`, `case "`+cmd+`"`+":"+"\n"+cmd+"."+cases.Title(language.Und, cases.NoLower).String(cmd)+"() \n"+"break \n"); err != nil {
		switchutility.Checklogger(err, "trying to update text in cmdrunner.go")
	}

	if err := switchutility.UpdateText("./switchtermer/cmd/cmdrunner/cmdrunner.go", `"github.com/golangast/switchterm`+dir+`/cmd/`+cmd+`"`, `// #import`, `"github.com/golangast/switchterm`+dir+`/cmd/`+cmd+`"`); err != nil {
		switchutility.Checklogger(err, "trying to update text in cmdrunner.go")
	}

}
