package add

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/cmd/cmdcreator"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switch/switchutility"
)

func Add() {

	listbash := []string{"bash", "custom"}

	//print directions
	switchutility.Directions()

	answerbash := switchselector.DigSingle(listbash, 1, "green", "red")

	switch answerbash {

	case "bash":
		fmt.Println("add a commnd..")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputcmd := scanner.Text()
		stripcmd := strings.TrimSpace(inputcmd)

		fmt.Println("add a description..")
		scannerdesc := bufio.NewScanner(os.Stdin)
		scannerdesc.Scan()
		inputnote := scannerdesc.Text()

		fmt.Println("add a tag..")
		scannertag := bufio.NewScanner(os.Stdin)
		scannertag.Scan()
		inputtag := scannertag.Text()
		striptag := strings.TrimSpace(inputtag)

		fmt.Println("what is the directory of bashfile?")
		scannerbash := bufio.NewScanner(os.Stdin)
		scannerbash.Scan()
		inputbashfile := scannerbash.Text()
		stripbash := strings.TrimSpace(inputbashfile)

		fmt.Println("what is the name of bashfile?")
		scannerbashname := bufio.NewScanner(os.Stdin)
		scannerbashname.Scan()
		inputbashfilename := scannerbashname.Text()
		stripbashname := strings.TrimSpace(inputbashfilename)

		bash := "true"

		tags.Create(stripcmd, inputnote, striptag, bash, stripbash, stripbashname)
		cmdcreator.CreateBashFile(stripbash, stripbashname)

	case "custom":
		fmt.Println("add a commnd..")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputcmd := scanner.Text()
		stripcmd := strings.TrimSpace(inputcmd)

		fmt.Println("add a description..")
		scannerdesc := bufio.NewScanner(os.Stdin)
		scannerdesc.Scan()
		inputdesc := scannerdesc.Text()

		fmt.Println("add a tag..")
		scannertag := bufio.NewScanner(os.Stdin)
		scannertag.Scan()
		inputtag := scannertag.Text()
		striptag := strings.TrimSpace(inputtag)
		bash := "false"

		tags.Create(stripcmd, inputdesc, striptag, bash, "", "")

		cmdcreator.Cmdcreator(inputcmd)

	}
}
