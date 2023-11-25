package switchutility

import (
	"fmt"
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func diff(a, b []string) []string {
	temp := map[string]int{}
	for _, s := range a {
		temp[s]++
	}
	for _, s := range b {
		temp[s]--
	}

	var result []string
	for s, v := range temp {
		if v != 0 {
			result = append(result, s)
		}
	}
	return result
}
