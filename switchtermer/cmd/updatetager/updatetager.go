package updatetager

import (
	"bufio"
	"fmt"
	"os"

	"github.com/golangast/switchterm/db/sqlite/tags"
)

func UpdateTager(oldtag string) error {

	fmt.Println("What do you want to name the tag?")

	scannerdesc := bufio.NewScanner(os.Stdin)
	scannerdesc.Scan()
	tagname := scannerdesc.Text()

	tags.UpdateTag(oldtag, tagname)

	return nil

}
