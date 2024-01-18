package settings

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golangast/switchterm/db/sqlite/sqlsettings"
)

func Settings() {
	fmt.Println("What directory do you want to use to store the commands?")
	scannerdesc := bufio.NewScanner(os.Stdin)
	scannerdesc.Scan()
	dir := scannerdesc.Text()
	stripdir := strings.TrimSpace(dir)

	sqlsettings.UpdateSettings(stripdir)

	Makefolder("./switchterm/" + dir)

}

// make any folder
func Makefolder(p string) error {
	if err := os.MkdirAll(p, os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("~~~~could not create"+p, err)
		return err
	}
	return nil
}
