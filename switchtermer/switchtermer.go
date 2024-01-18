package switchtermer

import (
	"fmt"
	"log/slog"

	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/cmd/add"
	"github.com/golangast/switchterm/switchtermer/cmd/search"
	"github.com/golangast/switchterm/switchtermer/data"
	"github.com/golangast/switchterm/switchtermer/settings"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switch/switchutility"
	"github.com/golangast/switchterm/switchtermer/window"
)

// takes in a list of strings and then lets you select search or select
// The `func` keyword is used to define a function in Go. In this code, it is used to define several
// functions, such as `SwitchCol`, `SwitchCall`, and others. These functions perform various tasks,
// such as switching between different actions based on user input, searching or selecting commands,
// adding commands, and displaying windows.
func SwitchCol(list []string, cols int, background, foreground string) []string {

	var atline int       //to know what line you are on
	var results []string // append to results

	//commands available
	lists := []string{"where to begin?", "settings", "search", "select", "add", "window", "data"}

	answer := switchselector.Menu(lists, 1, "purple", "purple")

	switch answer {
	case "where to begin?":
		colortermer.ColorizeOutPut("purple", "purple", "(start = initialize switchterm) - (settings = where directory of commands will be) - (search = search commands) \n (select = run commands)  - (add = add a command) - (window = group of tags) \n")
		fmt.Println("\n")
		SwitchCall()
	case "search":
		return search.Search(background, foreground, list, cols, atline)
	case "select":
		ans := switchselector.Dig(list, cols, background, foreground)
		return ans
	case "add":
		add.Add()
	case "window":
		window.Window()
	case "data":
		data.Data()
	case "settings":
		settings.Settings()
	default:
		return results
	}

	return results

}

func SwitchCall() {
	logger := loggers.CreateLogger()

	tt, err := tags.GetAll()
	if err != nil {
		logger.Error(
			"getting all tags",
			slog.String("error: ", err.Error()),
		)
	}
	var CMDS []string
	//turn into []string for the selector
	for _, item := range tt {
		CMDS = append(CMDS, item.CMD)
	}

	//function to search or select a command
	answ := SwitchCol(CMDS, 6, "magenta", "purple")
	answ = switchutility.RemoveDuplicateStr(answ)
	if len(answ) > 1 {

		//get notes from selection
		ta, err := tags.GetNoteByChosen(answ)
		if err != nil {
			logger.Error(
				"getting all tags by note",
				slog.String("error: ", err.Error()),
			)
		}
		//Print them
		colortermer.ColorizeOutPut("purple", "purple", "......NOTES.......\n")
		for _, v := range ta {
			colortermer.ColorizeOutPut("purple", "purple", "{ "+v.CMD+" ~")
			colortermer.ColorizeOutPut("dpurple", "bpurple", "Notes: "+v.Note+" }")
			fmt.Println("\n")

		}
	}
}
