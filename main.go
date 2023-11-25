package main

import (
	"fmt"

	"github.com/golangast/switchterm/switchtermer"
)

func main() {

	list := []string{
		"one", "two", "three", "four", "five", "six",
		"seven", "eight", "nine", "ten", "eleven", "twelve",
		"thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen",
		"nineteen", "twenty", "twenty-one", "twenty-two", "twenty-three", "twenty-four",
	}
	answ := switchtermer.SwitchCol(list, 6, "red", "green")
	fmt.Println(answ)

}
