package cmdcreator

import (
	"log/slog"

	ff "github.com/golangast/gentil/utility/ff"
	temps "github.com/golangast/gentil/utility/temp"
	text "github.com/golangast/gentil/utility/text"

	"github.com/golangast/switchterm/db/sqlite/sqlsettings"
	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/cmd/cmdcreator/temp"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Cmdcreator(cmd string) {

	var err error
	logger := loggers.CreateLogger()

	dir, err := sqlsettings.GetDir()
	if err != nil {
		logger.Error(
			"getting the directory",
			slog.String("error: ", err.Error()),
		)
	}

	//make file
	cmdfile, err := ff.Filefolder("./cmd/"+cmd, cmd+".go")
	if err != nil {
		logger.Error(
			"trying to create handler file",
			slog.String("error: ", err.Error()),
		)
	}
	m := make(map[string]string)
	m["cmd"] = cmd
	m["CMD"] = cases.Title(language.Und, cases.NoLower).String(cmd)
	//write to file
	err = temps.Writetemplate(temp.CmdTemp, cmdfile, m)
	if err != nil {
		logger.Error(
			"trying to update router.html",
			slog.String("error: ", err.Error()),
		)
	}

	//replace imports
	found := text.FindTextNReturn("./switchtermer/cmdrunner/cmdrunner.go", `//#addcmd`)

	if found != `//#addcmd` {
		err := text.UpdateText("./switchtermer/cmdrunner/cmdrunner.go", `//#addcmd`, `case "`+cmd+`"`+":"+"\n"+cmd+"."+cases.Title(language.Und, cases.NoLower).String(cmd)+"() \n"+"break \n"+`//#addcmd`)
		if err != nil {
			logger.Error(
				"trying to update text in cmdrunner.go",
				slog.String("error: ", err.Error()),
			)
		}
	} else {
		err := text.UpdateText("./switchtermer/cmdrunner/cmdrunner.go", `//#addcmd`, `case "`+cmd+`"`+":"+"\n"+cmd+"."+cases.Title(language.Und, cases.NoLower).String(cmd)+"() \n"+"break \n"+`//#addcmd`)
		if err != nil {
			logger.Error(
				"trying to update text in cmdrunner.go",
				slog.String("error: ", err.Error()),
			)
		}
	}

	foundimport := text.FindTextNReturn("./switchtermer/cmdrunner/cmdrunner.go", `// #import`)
	if foundimport == `// #import` {
		err := text.UpdateText("./switchtermer/cmdrunner/cmdrunner.go", `// #import`, `"github.com/golangast/switchterm`+dir+`/cmd/`+cmd+`"`+"\n"+`// #import`)
		if err != nil {
			logger.Error(
				"trying to update text in cmdrunner.go",
				slog.String("error: ", err.Error()),
			)
		}
	} else {
		err := text.UpdateText("./switchtermer/cmdrunner/cmdrunner.go", `// #import`, `"github.com/golangast/switchterm`+dir+`/cmd/`+cmd+`"`+"\n"+`// #import`)
		if err != nil {
			logger.Error(
				"trying to update text in cmdrunner.go",
				slog.String("error: ", err.Error()),
			)
		}
	}

}

func CreateBashFile(dir, file string) {
	logger := loggers.CreateLogger()

	// make file
	_, err := ff.Filefolder("."+dir, file+".bash")
	if err != nil {
		logger.Error(
			"trying to create bash file",
			slog.String("error: ", err.Error()),
		)
	}

}
