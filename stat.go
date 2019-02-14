// [x] calculate the min of temperature, pressure and wind speed
// [x] function that sorts the temp, presss, and wind speed
// [x] input - sorted numbers, output - median
// [ ] create a git repo
// [ ] tag the lessons via `git tag`
// [ ] put stat as a subpackage, without main
// [ ] update u-tests in different package (stat_test)
// [ ] make the functions exportable
// [ ] remove stat funcs from main package
// [ ] think about how you can make readData into 2 separate packages

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var (
	temp     []string
	pressure []string
	windSpd  []string
)

// readData reads the data from csv file and returns rows
func readData(name string) [][]string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)

	r.Comma = '\t'

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}

	return rows
}

// min returns minimum value from the dataset
func min(data []string) string {

	minData, _ := strconv.ParseFloat(data[0], 64)
	for _, v := range data {
		p, _ := strconv.ParseFloat(v, 64)
		if p < minData {
			minData = p
		}
	}
	return fmt.Sprintf("%.2f", minData)
}

// sortedDataSet  returns sorted dataset
func sortedDataSet(data []string) []string {

	floatData := []float64{}
	sortedData := []string{}

	// converting string dataset to float64 dataset
	for _, v := range data {
		p, _ := strconv.ParseFloat(v, 64)
		floatData = append(floatData, p)
	}
	// actual sorting
	sort.Float64s(floatData)

	// converting a float64 dataset to a string dataset
	for _, v := range floatData {
		p := strconv.FormatFloat(v, 'f', 2, 64)
		sortedData = append(sortedData, p)
	}

	return sortedData
}

// median gets the median number from a slice of numbers
// {3, 1, 4, 1}   returns 2
// {3, 1, 4, 1, 5} returns 3
func median(data []string) string {
	floatData := []float64{}
	var result float64

	// converting string dataset into float64 dataset
	for _, v := range data {
		p, _ := strconv.ParseFloat(v, 64)
		floatData = append(floatData, p)
	}

	half := len(floatData) / 2
	result = floatData[half]
	if len(floatData)%2 == 0 {
		result = (result + floatData[half-1]) / 2
	}

	// converting a result to string
	med := strconv.FormatFloat(result, 'f', 2, 64)

	return med
}

func main() {
	rows := readData("Environmental_Data_Deep_Moor_2015.csv")

	//fmt.Println(rows)
	for i := range rows {
		// pass the header of the file
		if i == 0 {
			continue
		}

		temp = append(temp, rows[i][1])
		pressure = append(pressure, rows[i][2])
		windSpd = append(windSpd, rows[i][7])
	}

	minTemp := min(temp)
	minPressure := min(pressure)
	minWindSpd := min(windSpd)
	sortedDataTemp := sortedDataSet(temp)
	sortedDataPress := sortedDataSet(pressure)
	sortedDataWindSpd := sortedDataSet(windSpd)

	fmt.Println("=====================")
	fmt.Println("Minimum Temperature: \t", minTemp)
	fmt.Println("Minimum Pressure: \t", minPressure)
	fmt.Println("Minimum Windspeed: \t", minWindSpd)
	fmt.Println("=====================")
	fmt.Println("Median of Temperature: \t", median(sortedDataTemp))
	fmt.Println("Median of Pressure: \t", median(sortedDataPress))
	fmt.Println("Median of Wind Speed: \t", median(sortedDataWindSpd))
}
