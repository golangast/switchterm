package switchselector

import (
	"fmt"
	"log/slog"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/cmd/cmdrunner"
	"github.com/golangast/switchterm/switchtermer/cmd/updatetager"
	"github.com/golangast/switchterm/switchtermer/switch/colortermer"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func Dig(list []string, cols int, background, foreground string) []string {
	logger := loggers.CreateLogger()
	var (
		atline int
		chosen []string
		remove bool
		exes   bool
	)
	//print in colunns
	switchutility.PrintColumns(cols, atline, list, chosen, background, foreground)
	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		//press arrows to change index to highlight selected item
		switch key.String() {
		case "up": //up arrow
			atlines, run, err := switchutility.UP(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "down": //down arrow
			atlines, run, err := switchutility.Down(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "right": //right arrow
			atlines, run, err := switchutility.Right(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "left": //left arrow
			atlines, run, err := switchutility.Left(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "c": //choose another
			chosen = append(chosen, list[atline])
			remove = false
			return false, nil
		case "r": //removing selection
			remove = true
			chosen = append(chosen, list[atline])
			return false, nil
		case "u": //removing selection
			remove = false
			exes = false
			err := updatetager.UpdateTager(list[atline])
			return false, err
		case "enter": //select and run
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
		logger.Error(
			"pressing keys",
			slog.String("error: ", err.Error()),
		)
	}
	//remove item after one has been chosen
	remove = switchutility.RemoveItemWChosen(remove, list, chosen) //it is this way because you cannot call keyboard.Listen in itself

	if exes {
		cmdrunner.CmdRunner(chosen)
	}

	return chosen
}

func DigSingle(list []string, cols int, background, foreground string) string {
	logger := loggers.CreateLogger()

	var (
		atline int
		chosen []string
		ans    string
	)

	switchutility.PrintColumnsWChosen(cols, atline, list, background, foreground)
	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		//press arrows to change index to highlight selected item
		switch key.String() {
		case "up": //up arrow
			atlines, run, err := switchutility.UP(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "down": //down arrow
			atlines, run, err := switchutility.Down(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "right": //left arrow
			atlines, run, err := switchutility.Right(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "left": //left arrow
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
		logger.Error(
			"pressing keys",
			slog.String("error: ", err.Error()),
		)
	}
	return ans

}

func Menu(list []string, cols int, background, foreground string) string {
	logger := loggers.CreateLogger()

	var (
		atline int
		chosen []string
		ans    string
	)

	fmt.Print("\033[H\033[2J")
	colortermer.ColorizeCol("purple", "purple", "(q-quit) - (c-multiselection) - (r-remove) - (enter-select/execute) - (u-update tag) - down/up/left/right")
	fmt.Println("\n")
	switchutility.PrintColumns(cols, atline, list, chosen, background, foreground)

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		//press arrows to change index to highlight selected item
		switch key.String() {
		case "up": //up arrow
			atlines, run, err := switchutility.UP(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "down": //down arrow
			atlines, run, err := switchutility.Down(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "right": //left arrow
			atlines, run, err := switchutility.Right(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "left": //left arrow
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
		logger.Error(
			"pressing keys",
			slog.String("error: ", err.Error()),
		)
	}
	return ans

}

func MenuInstuctions(list []string, cols int, background, foreground, instructions string) string {
	logger := loggers.CreateLogger()

	var (
		atline int
		chosen []string
		ans    string
	)

	fmt.Print("\033[H\033[2J")
	colortermer.ColorizeCol("purple", "purple", "(q-quit) - (c-multiselection) - (r-remove) - (enter-select/execute) - (u-update tag) - down/up/left/right")
	fmt.Println("\n")
	colortermer.ColorizeCol("purple", "purple", instructions)
	fmt.Println("\n")
	switchutility.PrintColumns(cols, atline, list, chosen, background, foreground)

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		//press arrows to change index to highlight selected item
		switch key.String() {
		case "up": //up arrow
			atlines, run, err := switchutility.UP(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "down": //down arrow
			atlines, run, err := switchutility.Down(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "right": //left arrow
			atlines, run, err := switchutility.Right(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "left": //left arrow
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
		logger.Error(
			"pressing keys",
			slog.String("error: ", err.Error()),
		)
	}
	return ans

}
