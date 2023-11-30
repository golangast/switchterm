package switchutility

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"slices"

	"github.com/golangast/switchterm/switchtermer/colortermer"
)

func UP(atline, rows, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	if atline >= 1 {
		atline--
	}
	PrintColumns(rows, cols, atline, list, chosen, background, foreground)

	return atline, false, nil
}

func Down(linecount, atline, rows, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	if atline <= linecount-2 {
		atline++
	}
	PrintColumns(rows, cols, atline, list, chosen, background, foreground)

	return atline, false, nil

}
func Right(linecount, atline, rows, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	if atline <= linecount-rows {
		atline = atline + rows
	}

	PrintColumns(rows, cols, atline, list, chosen, background, foreground)

	return atline, false, nil

}

func Left(linecount, atline, rows, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	if atline >= rows {
		atline = atline - rows
	}
	PrintColumns(rows, cols, atline, list, chosen, background, foreground)

	return atline, false, nil

}
func ClearDirections() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("q-quit|e-multiselection|enter-select|down/up/left/right|")
}

func Directions() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("")
}

func PrintColumns(rows, cols, atline int, list, chosen []string, background, foreground string) {

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

func PrintColumnsWChosen(rows, cols, atline int, list []string, background, foreground string) {
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
