package main

import (
	"fmt"
	"strconv"
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

	fmt.Println(dataSlice)
}
