package utils

import (
	"math"

	"github.com/guptarohit/asciigraph"
)

var siChars = []string{"", "k", "M", "G", "T", "P", "E", "Z", "Y"}

func getSi(num float64) (float64, string) {
	var n int
	for num >= 1000 && n != 7 {
		num = num / 1000
		n++
	}

	return math.Pow(1000, float64(n)), siChars[n]
}

// SiSlice applies si prefix to the slice
func SiSlice(slice []float64) ([]float64, string) {
	var sum float64
	var sLen float64
	for _, integer := range slice {
		sum += integer
		sLen++
	}

	div, siChar := getSi(sum / sLen)
	var newSlice []float64
	for _, integer := range slice {
		newSlice = append(newSlice, integer/div)
	}

	return newSlice, siChar
}

// Chart converts series to the asciigraph chart
func Chart(series []float64, caption string, unit string) string {
	newSeries, siChar := SiSlice(series)
	return asciigraph.Plot(newSeries, asciigraph.Height(5), asciigraph.Caption(caption+" ("+siChar+")"))
}
