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

	case "green":
		fcolor = 32

	default:
		fcolor = 0
	}

	switch BackgroundColor {
	case "black":
		bcolor = "40"
	case "red":
		bcolor = "41"

	case "green":
		bcolor = "42"

	default:
		bcolor = "9"
	}
	fmt.Printf("%s\x1b[37;"+bcolor+";%dm%-0s%s\x1b[37;9;%dm %s", "[ ", fcolor, line, "", 0, "]")
	fmt.Printf("\x1b[37;9;%dm %s", 0, "   ")

}
