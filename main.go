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

	provider_ := provider.NewFileProvider("/Users/milliyang/go/src/quant/spider/export")
	quant.RegisterProvider(provider_)
	quant.Run()

	if true {
		return
	}

	fSeries := series.NewFloatSeries()

	sma := indicator.NewSMA(fSeries, 6)

	smaOfSma := indicator.NewSMA(sma, 6)

	now := time.Now()
	var i64 float64
	for i := 0; i < 10; i++ {
		i64 = float64(i)
		now = now.Add(time.Second)
		fSeries.Append(&now, i64)
	}

	// fmt.Println(sma.Keys())
	fmt.Println("orignial:", fSeries.Values())
	fmt.Println("sma6:", sma.Values())
	fmt.Println("sma6 of sma6", smaOfSma.Values())

	// times := fSeries.Keys()
	// for _, one := range times {
	// 	fmt.Print(smaOfSma.ValueAtTime(&one), " ")
	// }
}
