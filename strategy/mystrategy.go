package strategy

import (
	"fmt"
	"quant/base/bar"
	"quant/base/strategy"
)

const (
	debug = true
)

func init() {
	if debug {
		fmt.Println("quanta.strategy package init")
	}
}

type MyStrategy struct {
	strategy.Strategy
}

func (this *MyStrategy) Init(symbol string) {
	if debug {
		fmt.Println("MyStrategy.Init()")
	}
	this.Strategy.Init(symbol)

	// ToDo: init your strategy below:
	this.Name = "MyStrategy"

}

func (this *MyStrategy) OnStrategyStart() {
	fmt.Println("MyStrategy", "OnStrategyStart", "Name", this.Name, "Symbol:", this.Symbol)
}

func (this *MyStrategy) OnStrategyStop() {
	fmt.Println("MyStrategy", "OnStrategyStop", "Name", this.Name, "Symbol:", this.Symbol)

}

func (this *MyStrategy) OnBarOpen(bar bar.Bar) {

}

func (this *MyStrategy) OnBar(bar bar.Bar) {
	fmt.Println("MyStrategy", this.Symbol, "OnBar")

}

func (this *MyStrategy) OnBarSlice(size int) {

}
