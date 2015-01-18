package main

import (
	"fmt"
	"quant"
	"quant/base/indicator"
	"time"

	"quant/provider"
	_ "quanta/project"
)

func main() {

	provider_ := provider.NewFileProvider("/Users/milliyang/go/src/quant/spider/export")
	quant.RegisterProvider(provider_)
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
