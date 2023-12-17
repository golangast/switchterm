package switchtermer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

// takes in a list of strings and then lets you select search or select
func SwitchCol(list []string, cols int, background, foreground string) []string {
	var atline int       //to know what line you are on
	var results []string // append to results

	// init map
	lines := make(map[int]string) // to cycle through lines

	// load values into map
	for i, item := range list {
		lines[i] = item
	}
	//commands available
	lists := []string{"search", "select", "add"}

	//print directions
	switchutility.Directions()

	answer := switchutility.DigSingle(lists, 1, "green", "red")

	switch answer {

	case "search":

		//commands available
		listsearch := []string{"cmd", "tag"}

		//print directions
		switchutility.Directions()

		//show commands to select.
		searchanswer := switchutility.DigSingle(listsearch, 1, "green", "red")

		switch searchanswer {
		case "cmd":
			//to capture the first two characters to search
			var letters string
			fmt.Println("type first letters you want to search by", "example: `th` and then press enter")
			n, err := fmt.Scanf("%s\n", &letters)
			if err != nil || n != 1 {
				fmt.Println(n, err)
			}

			if len(letters) > 1 {

				for _, s := range list {

					if strings.HasPrefix(s, letters) {
						results = append(results, s)
					}
				}

				if len(results) < 6 {
					cols = 1
				}
				// print in colunns
				switchutility.PrintColumnsWChosen(cols, atline, results, background, foreground)

				list = results
				answers := switchutility.Dig(list, cols, background, foreground)
				return answers
			} else {
				fmt.Println("choose another letter")
			}

		case "tag":
			var tagcmds []string
			var selectedtag []string
			//get all tags
			tagcmd, err := tags.GetAll()
			if err != nil {
				fmt.Println(err)
			}
			//collect the tags
			for _, v := range tagcmd {
				tagcmds = append(tagcmds, v.Tag)
			}
			//show them
			switchutility.PrintColumnsWChosen(cols, atline, tagcmds, background, foreground)
			//do a selection of tags
			answers := switchutility.DigSingle(tagcmds, cols, background, foreground)
			//get cmd by tag
			selectedtag, err = tags.GetCMDByTag(answers)
			if err != nil {
				fmt.Println(err)
			}
			//let the user select cmd from the cmds that were from the tags
			anstag := switchutility.Dig(selectedtag, cols, background, foreground)

			return anstag

		}

	case "select":
		answers := switchutility.Dig(list, cols, background, foreground)
		return answers
	case "add":
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
		tags.Create(inputcmd, inputdesc, inputtag)

	}

	return results
}
