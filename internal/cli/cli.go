package cli

import (
	"fmt"
	"strconv"
)

const (
	hello    = "Hi! Enter nums for calc metrics (write !end for break)..."
	error    = "Slice is empty."
	choice   = "Choice the mode:"
	modeInfo = "Mean\t(1)\nMedian\t(2)\nMode\t(3)\nSD\t(4)\nAll\t(5)\nRestart\t(!t)\nQuit\t(!q)\n"
)

func ScanData() []int {
	fmt.Println(hello)

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

	if len(dataSlice) == 0 {
		fmt.Println(error)
		return nil
	}
	return dataSlice
}
