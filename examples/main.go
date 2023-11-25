package main

import (
	"fmt"

	"github.com/golangast/switchterm/switchtermer"
)

func main() {
	//commands
	list := []string{
		"one", "two", "three", "four", "five", "six",
		"seven", "eight", "nine", "ten", "eleven", "twelve",
		"thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen",
		"nineteen", "twenty", "twenty-one", "twenty-two", "twenty-three", "twenty-four",
	}

	//function to search or select a command
	answ := switchtermer.SwitchCol(list, 6, "red", "green")

	//returned command
	fmt.Println("you have chosen the command: ", answ)

}
