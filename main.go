package main

import (
	"fmt"

	"github.com/golangast/switchterm/configure"

	"github.com/golangast/switchterm/switchtermer"
)

func main() {

	//configure.GenConfigure()
	// t := tags.Tags{}
	// tt, err := t.GetCMD("ls")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(tt)
	//commands
	cmds := configure.LoadConfig()

	//function to search or select a command
	answ := switchtermer.SwitchCol(cmds, 6, "red", "green")

	//returned command
	fmt.Println("you have chosen the command: ", answ)

}
