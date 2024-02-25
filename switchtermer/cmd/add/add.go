package add

import (
	"github.com/golangast/gentil/utility/ff"
	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/cmd/cmdcreator"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func Add() {

	listbash := []string{"bash", "custom"}

	//print directions
	switchutility.Directions()

	answerbash := switchselector.DigSingle(listbash, 1, "green", "red")

	switch answerbash {

	case "bash":

		stripcmd := switchutility.InputScanDirections("add a commnd..")
		inputnote := switchutility.InputScanDirections("add a description..")
		striptag := switchutility.InputScanDirections("add a tag..")
		stripbash := switchutility.InputScanDirections("what is the directory of bashfile?")
		stripbashname := switchutility.InputScanDirections("what is the name of bashfile?")

		bash := "true"
		if err := tags.Create(stripcmd, inputnote, striptag, bash, stripbash, stripbashname); err != nil {
			switchutility.Checklogger(err, "creating bash file")
		}

		if _, err := ff.Filefolder("."+stripbash, stripbashname+".bash"); err != nil {
			switchutility.Checklogger(err, "trying to create bash file")
		}

	case "custom":
		stripcmd := switchutility.InputScanDirections("add a commnd..")
		inputdesc := switchutility.InputScanDirections("add a description..")
		striptag := switchutility.InputScanDirections("add a tag..")

		bash := "false"

		if err := tags.Create(stripcmd, inputdesc, striptag, bash, "", ""); err != nil {
			switchutility.Checklogger(err, "creating bash file")
		}

		cmdcreator.Cmdcreator(stripcmd)

	}
}
