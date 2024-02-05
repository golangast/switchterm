package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/golangast/gentil/utility/ff"
	"github.com/golangast/switchterm/loggers"
	"github.com/golangast/switchterm/switchtermer"
	"github.com/golangast/switchterm/switchtermer/db/createtable"
)

func main() {
	logger := loggers.CreateLogger()

	if _, err := os.Stat("./db/data.db"); errors.Is(err, os.ErrNotExist) {
		_, err := ff.Makefile("./db/data.db")
		if err != nil {
			fmt.Print(err)
		}

		err = createtable.CreateTable()
		if err != nil {
			logger.Error(
				"creating table",
				slog.String("error: ", err.Error()),
			)
		}

	}

	switchtermer.SwitchCall()
}
