package project

import (
	"fmt"
	"quant"
	"quanta/strategy"
)

const (
	debug = true
)

func init() {
	if debug {
		fmt.Println("quanta.project.init")
	}

	proj := quant.NewProject()
	proj.Name = "ProjectA"

	proj.AddInstrument("s000001")
	proj.AddInstrument("s000002")
	proj.AddInstrument("s000003")
	proj.AddInstrument("s000004")
	proj.AddInstrument("s000005")
	proj.AddInstrument("s000006")
	proj.Strategy(strategy.MyStrategy{})
}
