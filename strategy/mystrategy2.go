package strategy

import (
	"fmt"
	"quant/base/bar"
	"quant/base/strategy"
)

func init() {
	if debug {
		fmt.Println("quanta.strategy package init")
	}
}

type MyStrategy2 struct {
	strategy.Strategy
}

func (this *MyStrategy2) Init(symbol string) {
	if debug {
		fmt.Println("MyStrategy2.Init()")
	}
	this.Strategy.Init(symbol)
	this.Name = "MyStrategy2"
}

func (this *MyStrategy2) OnStrategyStart() {
	fmt.Println("MyStrategy2", "OnStrategyStart", "Name", this.Name, "Symbol:", this.Symbol)
}

func (this *MyStrategy2) OnStrategyStop() {
	fmt.Println("MyStrategy2", "OnStrategyStop", "Name", this.Name, "Symbol:", this.Symbol)

}

func (this *MyStrategy2) OnBarOpen(bar bar.Bar) {

}

func (this *MyStrategy2) OnBar(bar bar.Bar) {
	fmt.Println("MyStrategy2", this.Symbol, "OnBar")

}

func (this *MyStrategy2) OnBarSlice(size int) {

}
