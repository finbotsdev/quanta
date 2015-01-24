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

	// proj.AddInstrument("STO000001")
	proj.AddInstrument("STO000002")
	// proj.AddInstrument("STO000012")
	// proj.AddInstrument("STO000004")
	// proj.AddInstrument("STO000005")
	// proj.AddInstrument("STO000006")
	proj.Strategy(strategy.MyStrategy{})
	// proj.Strategy(strategy.MyStrategy2{})

}
