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
	result.MEAN, _ = meanCalc(slice)
	result.MEDIAN, _ = medianCalc(slice)
	result.MODE, _ = modeCalc(slice)
	result.SD, _ = sdCalc(slice)
	return result, nil
}

func MakeCalc(m IMetrics, slice []int) (Metrics, error) {
	sort.Ints(slice)
	result, _ := m.Calculate(slice)
	return result, nil
}

func GetInfo(m Metrics) {
	fmt.Println("Mean:\t", m.MEAN)
	fmt.Println("Median:\t", m.MEDIAN)
	fmt.Println("Mode:\t", m.MODE)
	fmt.Printf("SD:\t %0.2f\n", m.SD)
}

func meanCalc(slice []int) (float32, error) {
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

func medianCalc(slice []int) (float32, error) {
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
		median = (float32(slice[index]) + float32(slice[index+1])) / 2
	}

	return median, nil
}

func sdCalc(slice []int) (float32, error) {
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

func modeCalc(slice []int) (int, error) {
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
