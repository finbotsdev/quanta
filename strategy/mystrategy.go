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
		fmt.Println("sample.quant.strategy.init")
	}
}

type MyStrategy struct {
	strategy.Strategy
}

func (this *MyStrategy) Init() {
	this.Strategy.Init()
	this.Name = "Strategy"
}

func (this *MyStrategy) OnStrategyStart() {
	fmt.Println("MyStrategy", "OnStrategyStart")
}

func (this *MyStrategy) OnStrategyStop() {

}

func (this *MyStrategy) OnBarOpen(bar bar.Bar) {

}

func (this *MyStrategy) OnBar(bar bar.Bar) {
	fmt.Println("MyStrategy", "OnBar")

}

func (this *MyStrategy) OnBarSlice(size int) {

}
