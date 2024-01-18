package switchutility

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"

	"github.com/golangast/switchterm/loggers"

	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/cmd/cmdremover"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
)

func UP(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	if atline >= 1 {
		atline--
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)
	return atline, false, nil
}

func Down(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	linecount := len(list)
	if atline <= linecount-2 {
		atline++
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)
	return atline, false, nil

}
func Right(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	linecount := len(list)
	rows := (len(list) + cols - 1) / cols
	if atline <= linecount-rows {
		atline = atline + rows
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)
	return atline, false, nil
}

func Left(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	rows := (len(list) + cols - 1) / cols
	if atline >= rows {
		atline = atline - rows
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)

	return atline, false, nil
}
func ClearDirections() {
	fmt.Print("\033[H\033[2J")
	colortermer.ColorizeCol("purple", "purple", "(q-quit) - (c-multiselection) - (r-remove) - (enter-select/execute) - (u-update tag) - down/up/left/right")
	fmt.Println("\n")

}

func Directions() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("")
}

func PrintColumns(cols, atline int, list, chosen []string, background, foreground string) {
	rows := (len(list) + cols - 1) / cols

	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {
			i := col*rows + row
			if i >= len(list) {
				break // This means the last column is not "full"
			}

			if i == atline {

				colortermer.ColorizeCol(background, foreground, list[atline])

			} else {
				if slices.Contains(chosen, list[i]) {
					colortermer.ColorizeCol("purple", foreground, list[i])

				} else {
					fmt.Printf("%-11s%s", list[i], " ")
				}
			}
		}
		fmt.Println() //yes this needs to be here for padding
	}
}

func PrintColumnsWChosen(cols, atline int, list []string, background, foreground string) {
	rows := (len(list) + cols - 1) / cols

	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {
			i := col*rows + row
			if i >= len(list) {
				break // This means the last column is not "full"
			}

			if i == atline {
				colortermer.ColorizeCol(background, foreground, list[atline])

			} else {
				fmt.Printf("%-11s%s", list[i], " ")

			}

		}
		fmt.Println() //yes this needs to be here for padding

	}
}

func RemoveItemWChosen(remove bool, list, chosen []string) bool {
	// if remove is true then remove the chosen
	if remove {

		//remove chosen from list
		for _, item := range chosen {
			index := slices.Index(list, item)
			if index > -1 {
				tags.DeleteTag(item)
				fmt.Println("removed: ", item)

			}

			cmdremover.RemoveCMD(item)
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

func BashFileOut(file string, commands []string) (string, string, error) {
	logger := loggers.CreateLogger()

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmds := strings.Join(commands, " ")
	f, err := os.Open(file)
	if err != nil {
		logger.Error(
			"opening bash file ",
			slog.String("error: "+file, err.Error()),
		)
	}
	defer f.Close()

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/dev/stdin", cmds)
	} else {
		cmd = exec.Command("bash", "/dev/stdin", cmds)
	}
	cmd.Stdin = f
	cmd.Stdout = os.Stdout

	if cmd.Err != nil {
		logger.Error(
			"trying to run bash command",
			slog.String("error: ", err.Error()),
		)
	}

	err = cmd.Run()
	if err != nil {
		logger.Error(
			"running bash command",
			slog.String("error: ", err.Error()),
		)
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	return stdout.String(), stderr.String(), nil
}

func RunApps(chosen []string) {
	logger := loggers.CreateLogger()

	t, err := tags.GetNoteByChosen(chosen)
	if err != nil {
		logger.Error(
			"getting all tags by note",
			slog.String("error: ", err.Error()),
		)
	}
	var argsset []string
	for _, v := range t {
		colortermer.ColorizeOutPut("dpurple", "purple", "What are the Args for "+v.CMD+"? (please use spaces) ||| Note: "+v.Note+"\n")
		scannerdesc := bufio.NewScanner(os.Stdin)
		scannerdesc.Scan()
		args := scannerdesc.Text()
		strArrayOne := strings.Split(args, " ")
		argsset = append(argsset, strArrayOne...)

		out, outerr, err := BashFileOut("."+v.Bashdir+"/"+v.Bashfile+".bash", argsset)
		if err != nil {
			logger.Error(
				"trying to execute bash command",
				slog.String("error: ", err.Error()),
			)
		}

		if out != "" {
			fmt.Println("\n")
			colortermer.ColorizeOutPut("dpurple", "bpurple", "output: "+out)
			fmt.Println("\n")
		}
		if outerr != "" {
			colortermer.ColorizeOutPut("dpurple", "bpurple", "output: "+outerr)
		}
	}

}

func Execute(file, types string, commands []string) (*exec.Cmd, error) {

	f, e := os.Open(file)
	if e != nil {
		log.Fatal(e.Error())
	}
	defer f.Close()
	cmd := exec.Command(types, commands...)
	cmd.Stdin = f
	cmd.Stdout = os.Stdout

	if cmd.Err != nil {
		return cmd, cmd.Err
	}

	if e := cmd.Run(); e != nil {
		log.Fatal(e.Error())
	}

	return cmd, nil
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
func RemoveDuplicateStr(strSlice []string) []string {
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

func Initialize() {

}
