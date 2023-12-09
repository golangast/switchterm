package switchtermer

import (
	"fmt"
	"slices"
	"strings"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/golangast/switchterm/configure"
	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

// takes in a list of strings and then lets you select search or select
func SwitchCol(list []string, cols int, background, foreground string) []string {
	var atline int       //to know what line you are on
	var results []string // append to results

	// init map
	lines := make(map[int]string) // to cycle through lines

	// load values into map
	for i, item := range list {
		lines[i] = item
	}
	//commands available
	lists := []string{"search", "select", "add"}

	//print directions
	switchutility.Directions()

	answer := DigSingle(lists, 1, "green", "red")

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
			switchutility.PrintColumnsWChosen(cols, atline, results, background, foreground)

			list = results
			answers := Dig(list, cols, background, foreground)
			return answers
		} else {
			fmt.Println("choose another letter")
		}
	case "select":
		answers := Dig(list, cols, background, foreground)
		return answers
	case "add":
		var cmd string
		fmt.Println("add a commnd..")
		_, err := fmt.Scanf("%s", &cmd)
		if err != nil {
			fmt.Println(err)
		}
		var note string
		fmt.Println("add a note..")
		_, err = fmt.Scanf("%s", &note)
		if err != nil {
			fmt.Println(err)
		}
		var tag string
		fmt.Println("add a tag..")
		_, err = fmt.Scanln(tag)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("add a fdsf..")

		_, err = fmt.Scanln(tag)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cmd, note, tag)
		tags.Create(cmd, note, tag)

	}

	return results
}

func Dig(list []string, cols int, background, foreground string) []string {
	var (
		atline int
		chosen []string
		remove bool
	)

	switchutility.ClearDirections()
	//print in colunns
	switchutility.PrintColumns(cols, atline, list, chosen, background, foreground)

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		//press arrows to change index to highlight selected item
		switch key.String() {
		case "up": //up arrow
			//print directions
			switchutility.ClearDirections()
			//make it select up
			atlines, run, err := switchutility.UP(atline, cols, background, foreground, list, chosen)
			//keep listening

			atline = atlines
			return run, err

		case "down": //down arrow
			//print directions
			switchutility.ClearDirections()
			//make it select down
			atlines, run, err := switchutility.Down(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err

		case "right": //left arrow
			//print directions
			switchutility.ClearDirections()
			//make it select right
			atlines, run, err := switchutility.Right(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err

		case "left": //left arrow
			//print directions
			switchutility.ClearDirections()
			//make it select left
			atlines, run, err := switchutility.Left(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err

		case "e": //choose another
			chosen = append(chosen, list[atline])
			fmt.Println("added: ", chosen)
			remove = false

			return false, nil // Return false to continue listening
		case "r": //removing selection
			remove = true
			chosen = append(chosen, list[atline])
			return false, nil // Return false to continue listening
		case "x": //removing selection
			chosen = append(chosen, list[atline])
			for _, v := range chosen {
				fmt.Println("running...: ", v)
				out, errout, err := switchutility.Shellout(v)
				if err != nil {
					fmt.Println("error: ", err)
				}
				if errout != "" {
					fmt.Println("error: ", errout)
				}
				fmt.Println("out: ", out)
			}
			fmt.Println("ran: ", chosen)
			remove = false

			return false, nil // Return false to continue listening
		case "enter": //enter
			chosen = append(chosen, list[atline])

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

	//if remove is true then remove the chosen
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
	return chosen

}

func DigSingle(list []string, cols int, background, foreground string) string {
	var (
		atline int
		chosen []string
		ans    string
	)

	switchutility.ClearDirections()
	switchutility.PrintColumnsWChosen(cols, atline, list, background, foreground)

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		//press arrows to change index to highlight selected item
		switch key.String() {
		case "up": //up arrow
			//print directions
			switchutility.ClearDirections()
			//make it select up
			atlines, run, err := switchutility.UP(atline, cols, background, foreground, list, chosen)
			//keep listening

			atline = atlines
			return run, err

		case "down": //down arrow
			//print directions
			switchutility.ClearDirections()
			//make it select down
			atlines, run, err := switchutility.Down(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err

		case "right": //left arrow
			//print directions
			switchutility.ClearDirections()
			//make it select right
			atlines, run, err := switchutility.Right(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err

		case "left": //left arrow
			//print directions
			switchutility.ClearDirections()
			//make it select left
			atlines, run, err := switchutility.Left(atline, cols, background, foreground, list, chosen)
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
