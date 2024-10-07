package main

import (
	"github.com/Saracomethstein/go_day_00/internal/anscombe"
	"github.com/Saracomethstein/go_day_00/internal/cli"
)

func main() {
	var choice string
	dataSlice := cli.ScanData()

	for {
		choice = cli.Menu()
		if choice == "!m" {
			continue
		}
		break
	}

	result, _ := anscombe.MakeCalc(anscombe.Metrics{}, dataSlice)
	anscombe.GetInfo(result, choice)
}
