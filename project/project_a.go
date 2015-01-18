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
		fmt.Println("sample.quant.project.init")
	}

	proj := quant.NewProject()
	proj.Name = "ProjectA"

	proj.Strategy(strategy.MyStrategy{})

}
