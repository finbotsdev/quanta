package main

import (
	"fmt"
	"quant"
	"quant/base/indicator"
	"time"

	_ "quanta/project"
)

func main() {

	quant.Run()

	sma := indicator.NewSMA(6)

	now := time.Now()
	var i64 float64
	for i := 0; i < 100; i++ {
		i64 = float64(i + 1)
		now = now.Add(time.Second)
		sma.Append(&now, i64)
	}
	// fmt.Println(sma.Keys())
	fmt.Println(sma.Values())
}
