package switchutility

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"slices"

	"github.com/golangast/switchterm/configure"
	"github.com/golangast/switchterm/switchtermer/colortermer"
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
	fmt.Println("q-quit|e-multiselection|x-execute|enter-select|down/up/left/right|")
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
					colortermer.ColorizeCol("red", foreground, list[i])

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

func RemoveItemWChosen(remove bool, list, chosen []string) {
	// if remove is true then remove the chosen
	if remove == true {
		//remove chosen from list
		for _, item := range chosen {
			index := slices.Index(list, item)
			if index > -1 {
				list = append(list[:index], list[index+1:]...)
			}
		}
		configure.RemoveCommand(list)
		fmt.Println("removed: ", list)
	}
	remove = false

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
