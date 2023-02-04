package main

import (
	"fmt"
	"flag"

	"github.com/alitoyn/bread-calculator/loaf"
)

var defaultFlour = 500.0

func main() {

	flourAmount := flag.Float64("flour", 0.0, "a float for flour")

	flag.Parse()

	if *flourAmount == 0.0 {
		fmt.Printf("--flour has not been set, defaulting to %0.f\n", defaultFlour)
		flourAmount = &defaultFlour
	}

	loaf := loaf.CalculteRatioFromFlour(*flourAmount, loaf.BasicLoafRatio)

	fmt.Printf("%+v\n", loaf)
}