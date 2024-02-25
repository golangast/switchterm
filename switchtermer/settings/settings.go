package settings

import (
	"github.com/golangast/switchterm/db/sqlite/sqlsettings"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func Settings() {
	stripdir := switchutility.InputScanDirections("What directory do you want to use to store the commands?")
	sqlsettings.UpdateSettings(stripdir)
	switchutility.Makefolder("./switchterm/" + stripdir)
}
