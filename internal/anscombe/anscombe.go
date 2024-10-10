package anscombe

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

type IMetrics interface {
	Calculate(slice []int) (Metrics, error)
}

type Metrics struct {
	MEAN   float32
	MEDIAN float32
	SD     float32
	MODE   int
}

func (m Metrics) Calculate(slice []int) (Metrics, error) {
	result := Metrics{}
	result.MEAN, _ = MeanCalc(slice)
	result.MEDIAN, _ = MedianCalc(slice)
	result.MODE, _ = ModeCalc(slice)
	result.SD, _ = SDCalc(slice)
	return result, nil
}

func MakeCalc(m IMetrics, slice []int) (Metrics, error) {
	if slice == nil {
		return Metrics{}, errors.New("slice is nil")
	}

	sort.Ints(slice)
	result, err := m.Calculate(slice)
	if err != nil {
		return Metrics{}, err
	}
	return result, nil
}

func GetInfo(m Metrics, choice string) {
	switch choice {
	case "1":
		fmt.Printf("Mean:\t %0.1f\n", m.MEAN)
	case "2":
		fmt.Printf("Median:\t %0.1f\n", m.MEDIAN)
	case "3":
		fmt.Println("Mode:\t", m.MODE)
	case "4":
		fmt.Printf("SD:\t %0.2f\n", m.SD)
	case "5":
		fmt.Printf("Mean:\t %0.1f\n", m.MEAN)
		fmt.Printf("Median:\t %0.1f\n", m.MEDIAN)
		fmt.Println("Mode:\t", m.MODE)
		fmt.Printf("SD:\t %0.2f\n", m.SD)
	}
}

func MeanCalc(slice []int) (float32, error) {
	err := checkEmpty(slice)
	if err != nil {
		return 0.0, err
	}

	var mean float32
	for i := range slice {
		mean += float32(slice[i])
	}
	return mean / float32(len(slice)), nil
}

func MedianCalc(slice []int) (float32, error) {
	err := checkEmpty(slice)
	if err != nil {
		return 0.0, err
	}

	lSlice := len(slice)
	var median float32
	var index int
	if lSlice%2 == 1 {
		index = lSlice / 2
		median = float32(slice[index])
	} else {
		index = lSlice / 2
		median = (float32(slice[index]) + float32(slice[index-1])) / 2
	}

	return median, nil
}

func SDCalc(slice []int) (float32, error) {
	err := checkEmpty(slice)
	if err != nil {
		return 0.0, err
	}

	var sum int
	for _, num := range slice {
		sum += num
	}
	mean := float64(sum) / float64(len(slice))

	var sumSquaredDiffs float64
	for _, num := range slice {
		diff := float64(num) - mean
		sumSquaredDiffs += diff * diff
	}

	variance := sumSquaredDiffs / float64(len(slice))
	standardDeviation := math.Sqrt(variance)

	return float32(standardDeviation), nil
}

func ModeCalc(slice []int) (int, error) {
	err := checkEmpty(slice)
	if err != nil {
		return 0, err
	}

	hMap := make(map[int]int)
	for _, num := range slice {
		hMap[num]++
	}

	maxFrequency := 0
	mode := slice[0]

	for num, count := range hMap {
		if count > maxFrequency || (count == maxFrequency && num < mode) {
			maxFrequency = count
			mode = num
		}
	}

	return mode, nil
}

func checkEmpty(slice []int) error {
	if len(slice) == 0 {
		return errors.New("slice is empty")
	}
	return nil
}
