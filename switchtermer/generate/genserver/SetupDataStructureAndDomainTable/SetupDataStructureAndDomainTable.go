package setupdatastructureanddomaintable

import (
	"github.com/golangast/switchterm/switchtermer/db/data"
	"github.com/golangast/switchterm/switchtermer/db/dbconn"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func SetupDataStructureAndDomainTable() {

	f := data.Fields{}

	f.Name = switchutility.InputScanDirections("Whats Name of your data structure?")
	datafields := switchutility.InputScanDirections("Whats the fields of the data? Please field.type then space then field.type like the example Name.string age.int")

	ff := switchutility.GetPropDatatype(datafields)
	f.Field = append(f.Field, ff...)

	if err := f.Create(); err != nil {
		switchutility.Checklogger(err, "store the data structure")
	}

	db, err := dbconn.DbConnection()
	switchutility.Checklogger(err, "opening up the connection to the database")

	if _, err := db.Query("CREATE TABLE IF NOT EXISTS `domains` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `domain` VARCHAR(64) NULL, `github` VARCHAR(255) NOT NULL)"); err != nil {
		switchutility.Checklogger(err, "store the data structure")
	}

	db.Close()

}
