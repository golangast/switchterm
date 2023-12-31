package window

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/db/sqlite/window"
	"github.com/golangast/switchterm/switchtermer/loggers"

	"github.com/golangast/switchterm/switchtermer/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func Window() {
	logger := loggers.CreateLogger()

	fmt.Println("Do you want to edit a window? or create a window?")

	listwindow := []string{"edit", "create"}

	// print directions
	switchutility.Directions()

	answer := switchselector.DigSingle(listwindow, 1, "green", "red")

	switch answer {
	case "edit":
		var names []string

		fmt.Println("Choose your window")
		windows, err := window.GetAll()
		if err != nil {
			logger.Error(
				"getting all windows",
				slog.String("error: ", err.Error()),
			)
		}
		for _, v := range windows {
			names = append(names, v.Name)
		}
		answer := switchselector.DigSingle(names, 1, "green", "red")

		fmt.Println("Choose a tag")

		tager, err := tags.GetAll()
		if err != nil {
			logger.Error(
				"getting all tags",
				slog.String("error: ", err.Error()),
			)
		}
		var tagnames []string
		for _, v := range tager {
			tagnames = append(tagnames, v.Tag)
		}

		answertag := switchselector.DigSingle(tagnames, 1, "green", "red")

		window.UpdateTag(answer, answertag)

	case "create":
		fmt.Println("What do you want to call this window?")
		scannerdesc := bufio.NewScanner(os.Stdin)
		scannerdesc.Scan()
		windowname := scannerdesc.Text()

		//print directions
		switchutility.Directions()

		fmt.Println("What tags do you want tied to this window?")

		tager, err := tags.GetAll()
		if err != nil {
			logger.Error(
				"getting all tags",
				slog.String("error: ", err.Error()),
			)
		}
		var tagss []string
		for _, v := range tager {
			tagss = append(tagss, v.Tag)
		}
		answer := switchselector.Dig(tagss, 1, "green", "red")
		answers := strings.Join(answer, " ")
		window.Create(windowname, answers)

	}
}
