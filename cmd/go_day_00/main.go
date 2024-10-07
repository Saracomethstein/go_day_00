package main

import (
	"github.com/Saracomethstein/go_day_00/internal/anscombe"
	"github.com/Saracomethstein/go_day_00/internal/cli"
)

func main() {
	dataSlice := cli.ScanData()
	result, _ := anscombe.MakeCalc(anscombe.Metrics{}, dataSlice)
	anscombe.GetInfo(result)
}
