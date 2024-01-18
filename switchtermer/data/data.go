package data

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/golangast/switchterm/db/sqlite/data"
	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
)

func Data() {

	//commands available
	lists := []string{"add", "delete", "update", "cmd", "tag"}

	answer := switchselector.Menu(lists, 1, "purple", "purple")

	switch answer {
	case "add":
		Add()
	case "delete":
		Delete()
	case "update":
		Update()
	}

	dd, err := data.GetDataByName("jim")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dd)
}

func GetTable(name string) data.TableData {
	tables := data.Getapptables()
	for _, v := range tables {
		for _, i := range v.Name {
			if i == "data" {
				return v
			}
		}

	}
	return tables[6]
}
func Update() {

	logger := loggers.CreateLogger()

	fmt.Println("What is the name of the data?")
	scannername := bufio.NewScanner(os.Stdin)
	scannername.Scan()
	inputname := scannername.Text()
	stripname := strings.TrimSpace(inputname)

	d, err := data.GetDataByName(stripname)
	if err != nil {
		logger.Error(
			"trying to connect to database",
			slog.String("error: ", err.Error()),
		)
	}

	table := GetTable("data")
	fmt.Println(table)
	fmt.Println(d)

	// fmt.Println("What field do you want to update?")
	// scannername := bufio.NewScanner(os.Stdin)
	// scannername.Scan()
	// inputname := scannername.Text()
	// stripname := strings.TrimSpace(inputname)
}
func Delete() {
	fmt.Println("What data do you want to delete?")
	scannerdelete := bufio.NewScanner(os.Stdin)
	scannerdelete.Scan()
	inputdelete := scannerdelete.Text()
	stripdelete := strings.TrimSpace(inputdelete)

	data.DeleteData(stripdelete)
}
func Add() {
	fmt.Println("What is the name of the data?")
	scannername := bufio.NewScanner(os.Stdin)
	scannername.Scan()
	inputname := scannername.Text()
	stripname := strings.TrimSpace(inputname)

	var fields []any

	// fields = append(fields, "")
	fields = append(fields, stripname)

	fmt.Println("How many fields are you wanting to create?")
	scannerfields := bufio.NewScanner(os.Stdin)
	scannerfields.Scan()
	inputfields := scannerfields.Text()
	stripfields := strings.TrimSpace(inputfields)

	ii, err := strconv.Atoi(stripfields)
	if err != nil {
		panic(err)
	}
	for i := 0; i < ii; i++ {
		fmt.Println("whats the next field?")
		scannerfield := bufio.NewScanner(os.Stdin)
		scannerfield.Scan()
		inputfield := scannerfield.Text()
		stripfield := strings.TrimSpace(inputfield)

		fields = append(fields, stripfield)

	}
	s := 13 - len(fields)
	for a := 0; a < s; a++ {
		fields = append(fields, " ")

	}

	data.Create(fields)
}
