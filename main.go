package main

import (
	"fmt"
	"quant"
	"quant/base/indicator"
	"quant/base/series"
	"time"

	"quant/provider"
	_ "quanta/project"
)

func main() {

	if false {
		provider_ := provider.NewFileProvider("/Users/milliyang/go/src/quant/spider/export")
		quant.RegisterProvider(provider_)
		quant.Run()
	}

	fSeries := series.NewFloatSeries()

	sma := indicator.NewSMA(fSeries, 6)

	// var iseries series.ISeries = sma

	now := time.Now()
	var i64 float64
	for i := 0; i < 10; i++ {
		i64 = float64(i + 1)

		now = now.Add(time.Second)
		fSeries.Append(&now, i64)
		// sma.Append(&now, i64)
		// iseries.Append(&now, i64)
	}

	// fmt.Println(sma.Keys())
	fmt.Println(fSeries.Values())
	fmt.Println(sma.Values())
}
