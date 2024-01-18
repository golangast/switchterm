package search

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switch/switchutility"
)

func Search(background, foreground string, list []string, cols, atline int) []string {
	var results []string
	logger := loggers.CreateLogger()

	// commands available
	listsearch := []string{"cmd", "tag"}

	// print directions
	switchutility.Directions()

	// show commands to select.
	searchanswer := switchselector.DigSingle(listsearch, 1, "green", "red")

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
			answers := switchselector.Dig(list, cols, background, foreground)
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
			logger.Error(
				"getting all tags",
				slog.String("error: ", err.Error()),
			)
		}
		//collect the tags
		for _, v := range tagcmd {
			tagcmds = append(tagcmds, v.Tag)
		}
		//show them
		switchutility.PrintColumnsWChosen(cols, atline, tagcmds, background, foreground)
		//do a selection of tags
		answers := switchselector.DigSingle(tagcmds, cols, background, foreground)
		//get cmd by tag
		selectedtag, err = tags.GetCMDByTag(answers)
		if err != nil {
			logger.Error(
				"getting all tags by cmd",
				slog.String("error: ", err.Error()),
			)
		}
		//let the user select cmd from the cmds that were from the tags
		anstag := switchselector.Dig(selectedtag, cols, background, foreground)

		return anstag

	}

	return results
}
