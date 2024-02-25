package cmdremover

import (
	"os"

	"github.com/golangast/switchterm/switchtermer/switchutility"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func RemoveCMD(cmd string) {

	if err := switchutility.UpdateText("./switchtermer/cmd/cmdrunner/cmdrunner.go", `"github.com/golangast/switchterm/cmd/`+cmd+`"`, "", ""); err != nil {
		switchutility.Checklogger(err, "trying to update text in cmdrunner.go")
	}

	if err := switchutility.UpdateText("./switchtermer/cmd/cmdrunner/cmdrunner.go", `case "`+cmd+`"`+":"+"\n"+cmd+"."+cases.Title(language.Und, cases.NoLower).String(cmd)+"()", "", ""); err != nil {
		switchutility.Checklogger(err, "trying to remove call")
	}

	if err := os.RemoveAll("./cmd/" + cmd); err != nil {
		switchutility.Checklogger(err, "trying to remove folder")
	}

}
