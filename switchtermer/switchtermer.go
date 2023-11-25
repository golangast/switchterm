package switchtermer

import (
	"fmt"
	"strings"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

// takes in a list of strings and then lets you select search or select
func SwitchCol(list []string, cols int, background, foreground string) []string {
	var atline int                        //to know what line you are on
	rows := (len(list) + cols - 1) / cols //rows
	var results []string                  // append to results

	// init map
	lines := make(map[int]string) // to cycle through lines

	// load values into map
	for i, item := range list {
		lines[i] = item
	}
	//commands available
	lists := []string{"search", "select"}

	//print directions
	switchutility.Directions()

	answer := DigSingle(lists, 1, "green", "red")
	fmt.Println("ans : ", answer)

	switch answer {
	case "search":

		var letters string
		fmt.Println("type first letters you want to search by", "example: `th` and then press enter")

		n, err := fmt.Scanf("%s\n", &letters)
		if err != nil || n != 1 {
			// handle invalid input
			fmt.Println(n, err)
		}
		fmt.Println(letters)
		//show what was pressed
		if len(letters) > 1 {

			for _, s := range list {

				if strings.HasPrefix(s, letters) {
					results = append(results, s)
				}
			}

			if len(results) < 6 {
				cols = 1
			}
			// print in colunns
			switchutility.PrintColumns(rows, cols, atline, results, background, foreground)

			list = results
			answers := Dig(list, cols, background, foreground)
			return answers
		} else {
			fmt.Println("choose another letter")
		}
	case "select":
		answers := Dig(list, cols, background, foreground)
		return answers
	}

	return results
}

func Dig(list []string, cols int, background, foreground string) []string {
	var atline int
	linecount := len(list)
	var ans []string
	rows := (len(list) + cols - 1) / cols
	//var results []string

	// init map
	lines := make(map[int]string)

	// load values into map
	for i, item := range list {
		lines[i] = item
	}
	switchutility.ClearDirections()
	//print in colunns
	switchutility.PrintColumns(rows, cols, atline, list, background, foreground)

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		//press arrows to change index to highlight selected item
		switch key.String() {
		case "up": //up arrow
			//print directions
			switchutility.ClearDirections()
			//make it select up
			atlines, run, err := switchutility.UP(atline, rows, cols, background, foreground, list)
			//keep listening

			atline = atlines
			return run, err

		case "down": //down arrow
			//print directions
			switchutility.ClearDirections()
			//make it select down
			atlines, run, err := switchutility.Down(linecount, atline, rows, cols, background, foreground, list)
			atline = atlines
			return run, err

		case "right": //left arrow
			//print directions
			switchutility.ClearDirections()
			//make it select right
			atlines, run, err := switchutility.Right(linecount, atline, rows, cols, background, foreground, list)
			atline = atlines
			return run, err

		case "left": //left arrow
			//print directions
			switchutility.ClearDirections()
			//make it select left
			atlines, run, err := switchutility.Left(linecount, atline, rows, cols, background, foreground, list)
			atline = atlines
			return run, err

		case "e": //choose another
			ans = append(ans, list[atline])
			fmt.Println("chosen: ", ans)
			return false, nil // Return false to continue listening

		case "enter": //enter
			ans = append(ans, list[atline])

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

func DigSingle(list []string, cols int, background, foreground string) string {
	var atline int
	linecount := len(list)
	var ans string
	rows := (len(list) + cols - 1) / cols
	//var results []string

	// init map
	lines := make(map[int]string)

	// load values into map
	for i, item := range list {
		lines[i] = item
	}
	switchutility.ClearDirections()
	//print in colunns
	switchutility.PrintColumns(rows, cols, atline, list, background, foreground)

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		//press arrows to change index to highlight selected item
		switch key.String() {
		case "up": //up arrow
			//print directions
			switchutility.ClearDirections()
			//make it select up
			atlines, run, err := switchutility.UP(atline, rows, cols, background, foreground, list)
			//keep listening

			atline = atlines
			return run, err

		case "down": //down arrow
			//print directions
			switchutility.ClearDirections()
			//make it select down
			atlines, run, err := switchutility.Down(linecount, atline, rows, cols, background, foreground, list)
			atline = atlines
			return run, err

		case "right": //left arrow
			//print directions
			switchutility.ClearDirections()
			//make it select right
			atlines, run, err := switchutility.Right(linecount, atline, rows, cols, background, foreground, list)
			atline = atlines
			return run, err

		case "left": //left arrow
			//print directions
			switchutility.ClearDirections()
			//make it select left
			atlines, run, err := switchutility.Left(linecount, atline, rows, cols, background, foreground, list)
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
