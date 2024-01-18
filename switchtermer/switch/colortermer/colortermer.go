package colortermer

import "fmt"

func ColorizeCol(ForegroundColor, BackgroundColor, line string) {
	var fcolor int
	var bcolor string
	switch ForegroundColor {
	case "black":
		fcolor = 30
	case "red":
		fcolor = 31

	case "purple":
		fcolor = 35
	case "dpurple":
		fcolor = 23
	default:
		fcolor = 0
	}

	switch BackgroundColor {
	case "black":
		bcolor = "40"
	case "red":
		bcolor = "41"

	case "purple":
		bcolor = "95"
	case "bpurple":
		bcolor = "91"
	case "magenta":
		bcolor = "45"
	default:
		bcolor = "9"
	}
	fmt.Printf("%s\x1b[37;"+bcolor+";%dm%-0s%s\x1b[37;9;%dm %s", "[ ", fcolor, line, "", 0, "]")
	fmt.Printf("\x1b[37;9;%dm %s", 0, "   ")

}

func ColorizeOutPut(ForegroundColor, BackgroundColor, line string) {
	var fcolor int
	var bcolor string
	switch ForegroundColor {
	case "black":
		fcolor = 30
	case "red":
		fcolor = 31

	case "purple":
		fcolor = 35
	case "dpurple":
		fcolor = 23
	default:
		fcolor = 0
	}

	switch BackgroundColor {
	case "black":
		bcolor = "40"
	case "red":
		bcolor = "41"

	case "purple":
		bcolor = "95"
	case "bpurple":
		bcolor = "91"
	case "magenta":
		bcolor = "45"
	default:
		bcolor = "9"
	}
	fmt.Printf("%s\x1b[37;"+bcolor+";%dm%-0s%s\x1b[37;9;%dm %s", " ", fcolor, line, "", 0, "")

}
