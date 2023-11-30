package main

import (
	"fmt"

	"github.com/golangast/switchterm/configure"
	"github.com/golangast/switchterm/switchtermer"
)

func main() {

	//commands
	cmds := configure.LoadConfig()

	configure.GenConfig()

	//function to search or select a command
	answ := switchtermer.SwitchCol(cmds, 6, "red", "green")

	//returned command
	fmt.Println("you have chosen the command: ", answ)

}
