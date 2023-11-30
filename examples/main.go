package main

import (
	"fmt"

	"github.com/golangast/switchterm/switchtermer"

	"github.com/golangast/switchterm/configure"
)

func main() {

	//commands
	cmds := configure.LoadConfig()
	//function to search or select a command
	answ := switchtermer.SwitchCol(cmds, 6, "red", "green")

	//returned command
	fmt.Println("you have chosen the command: ", answ)

}
