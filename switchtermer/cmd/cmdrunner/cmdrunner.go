package cmdrunner

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/golangast/switchterm/cmd/ff"
	"github.com/golangast/switchterm/cmd/server"
	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/switch/switchutility"
	// #import
)

func CmdRunner(chosen []string) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("")

	var err error
	logger := loggers.CreateLogger()
	slicetags, err := tags.GetAll()
	if err != nil {
		logger.Error(
			"trying to get all tags",
			slog.String("error: ", err.Error()),
		)
	}
	chosen = switchutility.RemoveDuplicateStr(chosen)
	for _, v := range slicetags {
		if slices.Contains(chosen, v.CMD) {
			if v.Bash != "true" {
				switch v.CMD {

				case "ff":
					ff.Ff()
					break
				case "server":
					server.Server()
					break
					//#addcmd

				default:

				}
				chosen = switchutility.Delete(chosen, v.CMD)
				chosen = switchutility.RemoveDuplicateStr(chosen)
			} else {

				if len(chosen) > 0 {
					switchutility.RunApps(chosen)
				}
			}
		}

	}

}
