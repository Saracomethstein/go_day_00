package main

import (
	"fmt"
	"strconv"

	"github.com/Saracomethstein/go_day_00/internal/anscombe"
)

func main() {
	var inputData string
	dataSlice := make([]int, 0)
	for {
		fmt.Scan(&inputData)
		if inputData == "!end" {
			break
		}

		num, err := strconv.Atoi(inputData)
		if err != nil {
			fmt.Println("Data is not a number!")
			continue
		}
		dataSlice = append(dataSlice, num)
	}

	result, _ := anscombe.MakeCalc(anscombe.Metrics{}, dataSlice)
	anscombe.GetInfo(result)
}
