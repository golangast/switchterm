package cmdrunner

import (
	"fmt"
	"slices"

	"github.com/golangast/switchterm/cmd/ff"
	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/switchutility"
	// #import
)

func CmdRunner(chosen []string) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("")

	var err error
	slicetags, err := tags.GetAll()
	switchutility.Checklogger(err, "get all tags")

	chosen = switchutility.RemoveDuplicateStr(chosen)
	for _, v := range slicetags {
		if slices.Contains(chosen, v.CMD) {
			if v.Bash != "true" {
				switch v.CMD {

				case "ff":
					ff.Ff()
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
