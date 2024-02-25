package updatetager

import (
	"github.com/golangast/switchterm/db/sqlite/tags"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func UpdateTager(oldtag string) error {

	tagname := switchutility.InputScanDirections("What do you want to name the tag?")

	tags.UpdateTag(oldtag, tagname)

	return nil

}
