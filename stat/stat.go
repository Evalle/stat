// [x] calculate the min of temperature, pressure and wind speed
// [x] function that sorts the temp, presss, and wind speed
// [x] input - sorted numbers, output - median
// [x] create a git repo
// [x] tag the lessons via `git tag`
// [x] put stat as a subpackage, without main
// [x] update u-tests in different package (stat_test)
// [x] make the functions exportable
// [x] remove stat funcs from main package
// [x] think about how you can make readData into 2 separate packages
// [x] TDD for errors
// [x] finish with space trimming
// [ ] http server in cmd

package stat

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// ReadData reads the data from csv file and returns rows

// Min returns minimum value from the dataset
func Min(data []string) (string, error) {

	if len(data) < 1 {
		return "", fmt.Errorf("the dataset should not be empty")
	}

	minData := math.MaxFloat64
	for _, v := range data {
		p, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			return "", err
		}
		if p < minData {
			minData = p
		}
	}
	return fmt.Sprintf("%.2f", minData), nil
}

// SortedDataSet  returns sorted dataset
func SortedDataSet(data []string) ([]string, error) {
	if len(data) < 1 {
		return nil, fmt.Errorf("the dataset should not be empty")
	}
	floatData := []float64{}
	sortedData := []string{}

	// converting string dataset to float64 dataset
	for _, v := range data {
		p, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			return nil, err
		}
		floatData = append(floatData, p)
	}
	// actual sorting
	sort.Float64s(floatData)

	// converting a float64 dataset to a string dataset
	for _, v := range floatData {
		p := strconv.FormatFloat(v, 'f', 2, 64)
		sortedData = append(sortedData, p)
	}

	return sortedData, nil
}

// Median gets the median number from a slice of numbers
// {3, 1, 4, 1}   returns 2
// {3, 1, 4, 1, 5} returns 3
func Median(data []string) (string, error) {
	if len(data) < 1 {
		return "", fmt.Errorf("the dataset should not be empty")
	}
	floatData := []float64{}
	var result float64

	// converting string dataset into float64 dataset
	for _, v := range data {
		p, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			return "", err
		}
		floatData = append(floatData, p)
	}

	half := len(floatData) / 2
	result = floatData[half]
	if len(floatData)%2 == 0 {
		result = (result + floatData[half-1]) / 2
	}

	// converting a result to string
	med := strconv.FormatFloat(result, 'f', 2, 64)

	return med, nil
}
