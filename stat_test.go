package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinData(t *testing.T) {
	table := []struct {
		expected string
		input    []string
	}{
		{
			"-1.00",
			[]string{"0.0", "-1.0", "19.00", "19.23"},
		},
		{
			"-1.10",
			[]string{"0.0", "-1.0", "-1.1", "19.23"},
		},
		{
			"19.00",
			[]string{"22.3", "22", "19.0", "19.23"},
		},
		{
			"-5.00",
			[]string{"0.0", "-5.0", "19.00", "-4.99"},
		},
	}

	for i, item := range table {
		t.Run(fmt.Sprintf("case #%d", i+1), func(t *testing.T) {
			assert.Equal(t, item.expected, min(item.input))
		})
	}
}

func TestSortedData(t *testing.T) {
	table := []struct {
		expected []string
		input    []string
	}{
		{
			[]string{"-1.00", "0.00", "19.00", "19.23"},
			[]string{"0.0", "-1.0", "19.00", "19.23"},
		},
		{
			[]string{"-1.10", "-1.00", "0.00", "19.23"},
			[]string{"0.0", "-1.0", "-1.1", "19.23"},
		},
		{
			[]string{"19.00", "19.23", "22.00", "22.30"},
			[]string{"22.3", "22", "19.0", "19.23"},
		},
		{
			[]string{"-5.00", "-4.99", "0.00", "19.00"},
			[]string{"0.0", "-5.0", "19.00", "-4.99"},
		},
	}

	for i, item := range table {
		t.Run(fmt.Sprintf("case #%d", i+1), func(t *testing.T) {
			assert.Equal(t, item.expected, sortedDataSet(item.input))
		})
	}
}

func TestMedian(t *testing.T) {

	table := []struct {
		expected string
		input    []string
	}{
		{
			"3.00",
			[]string{"1", "2", "3", "4", "5"},
		},
		{
			"-0.50",
			[]string{"-1.1", "-1.0", "0.0", "19.23"},
		},
		{
			"21.62",
			[]string{"19.0", "19.23", "24.0", "22.3"},
		},
		{
			"-1.50",
			[]string{"-5.0", "-4.99", "2.0", "19.00"},
		},
	}
	for i, item := range table {
		t.Run(fmt.Sprintf("case #%d", i+1), func(t *testing.T) {
			assert.Equal(t, item.expected, median(item.input))
		})
	}
}
