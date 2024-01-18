package cmdremover

import (
	"log/slog"
	"os"

	"github.com/golangast/switchterm/loggers"

	text "github.com/golangast/gentil/utility/text"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func RemoveCMD(cmd string) {
	logger := loggers.CreateLogger()
	foundimport := text.FindTextNReturn("./switchtermer/cmdrunner/cmdrunner.go", `"github.com/golangast/switchterm/cmd/`+cmd+`"`)
	if foundimport != `"github.com/golangast/switchterm/cmd/`+cmd+`"` {
		err := text.UpdateText("./switchtermer/cmdrunner/cmdrunner.go", `"github.com/golangast/switchterm/cmd/`+cmd+`"`, ``)
		if err != nil {
			logger.Error(
				"trying to remove import",
				slog.String("error: ", err.Error()),
			)
		}
	}

	found := text.FindTextNReturn("./switchtermer/cmdrunner/cmdrunner.go", `case "`+cmd+`"`+":"+"\n"+cmd+"."+cases.Title(language.Und, cases.NoLower).String(cmd)+"() \n"+"break")
	if found != `case "`+cmd+`"`+":"+"\n"+cmd+"."+cases.Title(language.Und, cases.NoLower).String(cmd)+"() \n"+"break" {
		err := text.UpdateText("./switchtermer/cmdrunner/cmdrunner.go", `case "`+cmd+`"`+":"+"\n"+cmd+"."+cases.Title(language.Und, cases.NoLower).String(cmd)+"() \n"+"break", ``)
		if err != nil {
			logger.Error(
				"trying to remove call",
				slog.String("error: ", err.Error()),
			)
		}
	}

	err := os.RemoveAll("./cmd/" + cmd)
	if err != nil {
		logger.Error(
			"trying to remove folder",
			slog.String("error: ", err.Error()),
		)
	}
}
