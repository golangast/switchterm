package main

import (
	"fmt"

	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer"
)

func main() {

	tt, err := tags.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	var CMDS []string
	//turn into []string for the selector
	for _, item := range tt {
		CMDS = append(CMDS, item.CMD)
	}
	//function to search or select a command
	answ := switchtermer.SwitchCol(CMDS, 6, "red", "green")

	//get notes from selection
	ta, err := tags.GetNoteByChosen(answ)
	if err != nil {
		fmt.Println(err)
	}
	//Print them
	for _, v := range ta {
		fmt.Println("CMD: ", v.CMD)
		fmt.Println("Notes: ", v.Note)

	}

}
