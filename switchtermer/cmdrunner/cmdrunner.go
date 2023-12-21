package cmdrunner

import (
	"bytes"
	"fmt"
	"log/slog"
	"os/exec"
	"runtime"
	"slices"

	"github.com/golangast/switchterm/cmd/aa"
	"github.com/golangast/switchterm/cmd/ee"
	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/loggers"
	// #import
)

func CmdRunner(exes bool, chosen []string) bool {
	if exes == true {
		var err error
		logger := loggers.CreateLogger()
		slicetags, err := tags.GetAll()
		if err != nil {
			logger.Error(
				"trying to get all tags",
				slog.String("error: ", err.Error()),
			)
		}
		for _, v := range slicetags {
			if slices.Contains(chosen, v.CMD) {
				//chosen = slices.Delete(chosen, 1, slices.Index(chosen, v.CMD))

				switch v.CMD {

				case "aa":
					aa.Aa()
					break
				case "ee":
					ee.Ee()
					break
					//#addcmd

				default:

				}
				i := slices.Index(chosen, v.CMD)
				chosen = append(chosen[:i], chosen[i+1:]...)
				if len(chosen) > 0 {
					RunApps(chosen)
				}
			}

		}
	}
	return false
}

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
func RunApps(chosen []string) {
	for _, v := range chosen {
		fmt.Println("running...: ", v)
		out, errout, err := Shellout(v)
		if err != nil {
			fmt.Println("error: ", err)
		}
		if errout != "" {
			fmt.Println("error: ", errout)
		}
		fmt.Println("out: ", out)
	}
	fmt.Println("ran: ", chosen)
}
