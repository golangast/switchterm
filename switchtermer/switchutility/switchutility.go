package switchutility

import (
	"fmt"
	"slices"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/cmdrunner"
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
	fmt.Println("q-quit|c-multiselection|x-select/execute||enter-select/execute|down/up/left/right|")
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

func RemoveItemWChosen(remove bool, list, chosen []string) bool {
	// if remove is true then remove the chosen
	if remove == true {
		//remove chosen from list
		for _, item := range chosen {
			index := slices.Index(list, item)
			if index > -1 {
				tags.DeleteTag(item)
				fmt.Println("removed: ", item)

			}
		}
	}

	return false

}

func Dig(list []string, cols int, background, foreground string) []string {
	var (
		atline int
		chosen []string
		remove bool
		exes   bool
	)
	ClearDirections()
	//print in colunns
	PrintColumns(cols, atline, list, chosen, background, foreground)
	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		//press arrows to change index to highlight selected item
		switch key.String() {
		case "up": //up arrow
			//make it select up
			atlines, run, err := UP(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "down": //down arrow
			//make it select down
			atlines, run, err := Down(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "right": //right arrow
			//make it select right
			atlines, run, err := Right(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "left": //left arrow
			//make it select left
			atlines, run, err := Left(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "c": //choose another
			chosen = append(chosen, list[atline])
			remove = false
			return false, nil // Return false to continue listening
		case "r": //removing selection
			remove = true
			chosen = append(chosen, list[atline])
			return false, nil // Return false to continue listening
		case "x":
			chosen = append(chosen, list[atline])
			cmdrunner.CmdRunner(exes, chosen) //runs the commands
			remove = false
			return false, nil // Return false to continue listening
		case "enter": //enter
			chosen = append(chosen, list[atline])
			exes = true
			return true, nil
		case "q", "esc", "ctrl+c": //to quit
			return true, nil
		default:
			fmt.Println(key.String())
			return false, nil // Return false to continue listening
		}

	})
	if err != nil {
		fmt.Println(err)
	}
	//remove item after one has been chosen
	remove = RemoveItemWChosen(remove, list, chosen) //it is this way because you cannot call keyboard.Listen in itself
	exes = cmdrunner.CmdRunner(exes, chosen)

	return chosen
}

func DigSingle(list []string, cols int, background, foreground string) string {
	var (
		atline int
		chosen []string
		ans    string
	)

	ClearDirections()
	PrintColumnsWChosen(cols, atline, list, background, foreground)
	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		//press arrows to change index to highlight selected item
		switch key.String() {
		case "up": //up arrow
			//make it select up
			atlines, run, err := UP(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err

		case "down": //down arrow
			//make it select down
			atlines, run, err := Down(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err

		case "right": //left arrow
			//make it select right
			atlines, run, err := Right(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err

		case "left": //left arrow
			//make it select left
			atlines, run, err := Left(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err

		case "enter": //enter
			ans = list[atline]
			return true, nil
		case "q", "esc", "c", "ctrl+c": //to quit
			return true, nil

		default:
			fmt.Println(key.String())
			return false, nil // Return false to continue listening
		}
	})

	if err != nil {
		fmt.Println(err)
	}
	return ans

}
