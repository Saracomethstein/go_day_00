package main

import (
	"fmt"
	"os"

	"github.com/Saracomethstein/go_day_00/internal/anscombe"
	"github.com/Saracomethstein/go_day_00/internal/cli"
)

func main() {
	var choice string
	dataSlice, _ := cli.ScanData()
	result, err := anscombe.MakeCalc(anscombe.Metrics{}, dataSlice)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	for {
		choice = cli.Menu()
		if choice == "!m" {
			continue
		}
		break
	}
	anscombe.GetInfo(result, choice)
}
