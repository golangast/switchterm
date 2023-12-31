package add

import (
	"bufio"
	"fmt"
	"os"

	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/cmdcreator"
	"github.com/golangast/switchterm/switchtermer/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
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

		fmt.Println("add a description..")
		scannerdesc := bufio.NewScanner(os.Stdin)
		scannerdesc.Scan()
		inputnote := scannerdesc.Text()

		fmt.Println("add a tag..")
		scannertag := bufio.NewScanner(os.Stdin)
		scannertag.Scan()
		inputtag := scannertag.Text()

		bash := "true"
		tags.Create(inputcmd, inputnote, inputtag, bash)

	case "custom":
		fmt.Println("add a commnd..")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputcmd := scanner.Text()

		fmt.Println("add a description..")
		scannerdesc := bufio.NewScanner(os.Stdin)
		scannerdesc.Scan()
		inputdesc := scannerdesc.Text()

		fmt.Println("add a tag..")
		scannertag := bufio.NewScanner(os.Stdin)
		scannertag.Scan()
		inputtag := scannertag.Text()

		fmt.Println("commands: ", inputcmd)
		fmt.Println("description: ", inputdesc)
		fmt.Println("tag: ", inputtag)
		bash := "false"

		tags.Create(inputcmd, inputdesc, inputtag, bash)

		cmdcreator.Cmdcreator(inputcmd)

	}
}
