package cli

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	hello     = "Hi! Enter nums for calc metrics (write !end for break)..."
	modeInfo  = "Choice the mode:\nMean\t(1)\nMedian\t(2)\nMode\t(3)\nSD\t(4)\nAll\t(5)\n"
	errorMenu = "Bad choice. Try again..."
)

func ScanData() ([]int, error) {
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
		return nil, errors.New("slice is empty")
	}
	return dataSlice, nil
}

func Menu() string {
	fmt.Printf("%s", modeInfo)

	var ch string
	fmt.Scan(&ch)

	switch ch {
	case "1", "2", "3", "4", "5":
		return ch
	default:
		fmt.Println(errorMenu)
		return "!m"
	}
}
