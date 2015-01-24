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

	if true {
		provider_ := provider.NewFileProvider("/Users/milliyang/go/src/quant/spider/export")
		quant.RegisterProvider(provider_)
		quant.Run()
		return
	}

	fSeries := series.NewFloatSeries()

	sma := indicator.NewSMA(fSeries, 16)

	smaOfSma := indicator.NewSMA(sma, 16)
	macd := indicator.NewMACD(fSeries, 16, 26, 9)

	now := time.Now()
	var i64 float64
	for i := 0; i < 100; i++ {
		i64 = float64(i)
		now = now.Add(time.Second)
		fSeries.Append(&now, i64)
	}

	// fmt.Println(sma.Keys())
	printslice("orignial:", fSeries.Values())
	printslice("sma6:", sma.Values())
	printslice("sma6 of sma6", smaOfSma.Values())

	fmt.Println("macd:")
	printslice("short ema:", macd.ShortEmaValues())
	printslice("long ema:", macd.LongEmaValues())
	printslice("diff:", macd.DiffValues())
	printslice("dem:", macd.DemValues())
	printslice("osc:", macd.OscValues())

	// times := fSeries.Keys()
	// for _, one := range times {
	// 	fmt.Print(smaOfSma.ValueAtTime(&one), " ")
	// }
}

func printslice(tag string, values []float64) {
	if len(tag) < 7 {
		fmt.Print(tag, "\t\t[")
	} else {
		fmt.Print(tag, "\t[")
	}
	count := 0
	for _, one := range values {
		count++
		if count < 90 {
			continue
		}
		fmt.Printf("%0.3f ", one)
	}
	fmt.Print("]\n")
}
