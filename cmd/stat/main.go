package main

import (
	"fmt"
	"log"

	"github.com/evalle/stat/data"
	"github.com/evalle/stat/formats"
	"github.com/evalle/stat/resources"

	"github.com/evalle/stat/stat"
)

type myFetchParser struct {
	resources.LocalFile
	formats.CSV
}

func extract(rows [][]string) (temp, pressure, windSpd []string) {
	temp, pressure, windSpd = []string{}, []string{}, []string{}

	for i := range rows {
		if i == 0 {
			continue
		}
		temp = append(temp, rows[i][1])
		pressure = append(pressure, rows[i][2])
		windSpd = append(windSpd, rows[i][7])
	}

	return
}

func min(temp, pressure, windSpd []string) (string, string, string) {
	return stat.Min(temp), stat.Min(pressure), stat.Min(windSpd)
}

func median(temp, pressure, windSpd []string) (string, string, string) {
	return stat.Median(stat.SortedDataSet(temp)),
		stat.Median(stat.SortedDataSet(pressure)),
		stat.Median(stat.SortedDataSet(windSpd))
}

func main() {
	rows, err := data.Read(
		&myFetchParser{
			LocalFile: resources.LocalFile{FileName: "Environmental_Data_Deep_Moor_2015.csv"},
			CSV:       formats.CSV{Separator: '\t'},
		})
	if err != nil {
		log.Fatal(err)
	}

	temp, pressure, windSpd := extract(rows)

	minTemp, minPressure, minWindSpd := min(temp, pressure, windSpd)

	medianTemp, medianPressure, medianWindSpd := median(temp, pressure, windSpd)

	fmt.Println("=====================")
	fmt.Println("Minimum Temperature: \t", minTemp)
	fmt.Println("Minimum Pressure: \t", minPressure)
	fmt.Println("Minimum Windspeed: \t", minWindSpd)
	fmt.Println("=====================")
	fmt.Println("Median of Temperature: \t", medianTemp)
	fmt.Println("Median of Pressure: \t", medianPressure)
	fmt.Println("Median of Wind Speed: \t", medianWindSpd)
}
