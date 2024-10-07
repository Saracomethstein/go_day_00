package anscombe

type IMetrics interface {
	Calculate([]int) (float32, error)
}

type Metrics struct {
	MEAN   float32
	MEDIAN float32
	SD     float32
	MODE   int
}

// Calculate metrics //
func (m Metrics) Calculate([]int) (float32, error) {
	var result float32

	return result, nil
}
