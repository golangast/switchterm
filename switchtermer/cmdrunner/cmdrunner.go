package cmdrunner

import (
	"bytes"
	"fmt"
	"log/slog"
	"os/exec"
	"runtime"
	"slices"

	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/loggers"

	"github.com/golangast/switchterm/cmd/aa"
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
				if v.Bash != "bash" {
					switch v.CMD {

					case "aa":
						aa.Aa()
						break

						//#addcmd

					default:

					}
					chosen = Delete(chosen, v.CMD)
					chosen = removeDuplicateStr(chosen)
				} else {

					if len(chosen) > 0 {
						RunApps(chosen)
					}
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
func Delete[T comparable](collection []T, el T) []T {
	idx := Find(collection, el)
	if idx > -1 {
		return slices.Delete(collection, idx, idx+1)
	}
	return collection
}

func Find[T comparable](collection []T, el T) int {
	for i := range collection {
		if collection[i] == el {
			return i
		}
	}
	return -1
}
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
