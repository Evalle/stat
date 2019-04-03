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
	statMinTemp, err := stat.Min(temp)
	if err != nil {
		log.Fatal(err)
	}

	statMinPressure, err := stat.Min(pressure)
	if err != nil {
		log.Fatal(err)
	}

	statMinWindSpd, err := stat.Min(windSpd)
	if err != nil {
		log.Fatal(err)
	}

	return statMinTemp, statMinPressure, statMinWindSpd
}

func median(temp, pressure, windSpd []string) (string, string, string) {
	//
	statSortedDataSetTemp, err := stat.SortedDataSet(temp)
	if err != nil {
		log.Fatal(err)
	}

	statMedianTemp, err := stat.Median(statSortedDataSetTemp)
	if err != nil {
		log.Fatal(err)
	}

	//
	statSortedDataSetPressure, err := stat.SortedDataSet(pressure)
	if err != nil {
		log.Fatal(err)
	}

	statMedianPressure, err := stat.Median(statSortedDataSetPressure)
	if err != nil {
		log.Fatal(err)
	}

	//
	statSortedDataSetWindSpd, err := stat.SortedDataSet(windSpd)
	if err != nil {
		log.Fatal(err)
	}

	statMedianWindSpd, err := stat.Median(statSortedDataSetWindSpd)
	if err != nil {
		log.Fatal(err)
	}

	return statMedianTemp, statMedianPressure, statMedianWindSpd
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
