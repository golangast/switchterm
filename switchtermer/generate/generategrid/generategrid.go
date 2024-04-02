package generategrid

import (
	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/db/grid"
	"github.com/golangast/switchterm/switchtermer/db/handler"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func Grid() {
	d, err := domain.GetStringDomains()
	switchutility.Checklogger(err, "getting all domains for grid")

	chosendomain := switchselector.MenuInstuctions(d, 1, "purple", "purple", "Which website are you going to run the database server for?")

	h, err := handler.GetStringHandlers()
	switchutility.Checklogger(err, "getting all handler for grid")

	gridhandler := switchselector.MenuInstuctions(h, 1, "purple", "purple", "Whats Name of the handler you want to use?")

	gridName := switchutility.InputScan("what is the name of the grid?")
	gridLayout := switchutility.InputScan("what is going to be the grid? (f=full, h=half, m=middler, q=quarter) example fhm")

	//store grid
	g := grid.Grid{Name: gridName, Domain: chosendomain, Handler: gridhandler, GridLayout: gridLayout}
	if ok, _ := g.Exists(g.Name); !ok {
		g.Create()
	}

	var gridstring string
	for _, l := range g.GridLayout {
		switch string(l) {
		case "f":
			gridstring += full
		case "h":
			gridstring += half
		case "m":
			gridstring += middler
		case "q":
			gridstring += quarters
		}
	}

	if err := switchutility.UpdateText(chosendomain+"/assets/templates/"+gridhandler+"/"+gridhandler+".html", `<!-- #grid -->`, `<!-- #grid -->`, gridstring+"\n"); err != nil {
		switchutility.Checklogger(err, "trying to update the grid in template file")
	}

}

var full = `<div class="full">full width</div>`
var half = `<div class="lhalf">left half</div><div class="rhalf">right half</div>`
var quarters = `<div class="quarter1">quarter1/div>
<div class="quarter2">quarter2</div>
<div class="quarter3">quarter3</div>
<div class="quarter4">quarter4</div>`

var middler = `<div class="middler">middler</div><div class="middler">middler</div>`
