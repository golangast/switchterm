package main

import (
	"fmt"

	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer"
)

func main() {

	//configure.GenConfigure()
	t := tags.Tags{}
	// tt, err := t.GetCMD("ls")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(tt)
	//commands
	// cmds := configure.LoadConfig()

	tt, err := t.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	var CMDS []string // to cycle through lines

	for _, item := range tt {
		CMDS = append(CMDS, item.CMD)
	}
	fmt.Println(CMDS)
	//function to search or select a command
	answ := switchtermer.SwitchCol(CMDS, 6, "red", "green")

	//returned command
	fmt.Println("you have chosen the command: ", answ)

}
