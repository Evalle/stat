package stat_test

import (
	"fmt"
	"testing"

	"github.com/evalle/stat/stat"
	"github.com/stretchr/testify/assert"
)

func TestMinData(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
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
			{
				"-5.00",
				[]string{" 0.0", "-5.0", "19.00", "-4.99"},
			},
		}

		for i, item := range table {
			t.Run(fmt.Sprintf("case #%d", i+1), func(t *testing.T) {
				value, err := stat.Min(item.input)
				assert.NoError(t, err, "there should be no error")
				assert.Equal(t, item.expected, value, "should be equal")
			})
		}
	})

	t.Run("Errors", func(t *testing.T) {
		table := []struct {
			expected string
			input    []string
		}{
			{
				"strconv.ParseFloat: parsing \"word\": invalid syntax",
				[]string{"word", "-1.0", "19.00", "19.23"},
			},
			{
				"strconv.ParseFloat: parsing \"0,0\": invalid syntax",
				[]string{"0,0", "-1.0", "-1.1", "19.23"},
			},
			{
				"the dataset should not be empty",
				[]string{},
			},
			{
				"the dataset should not be empty",
				nil,
			},
			{
				"strconv.ParseFloat: parsing \"-1,1\": invalid syntax",
				[]string{"0.0", "-1.0", "-1,1", "19.23"},
			},
		}

		for i, item := range table {
			t.Run(fmt.Sprintf("case #%d", i+1), func(t *testing.T) {
				value, err := stat.Min(item.input)
				assert.Empty(t, value, "value should be empty")
				assert.EqualError(t, err, item.expected, "should be error")
			})
		}
	})
}

func TestSortedData(t *testing.T) {

	t.Run("success", func(t *testing.T) {
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
			{
				[]string{"-5.00", "-4.99", "0.00", "19.00"},
				[]string{" 0.0", "-5.0", "19.00", "-4.99"},
			},
		}
		for i, item := range table {
			t.Run(fmt.Sprintf("case #%d", i+1), func(t *testing.T) {
				val, err := stat.SortedDataSet(item.input)
				assert.NoError(t, err, "should not be error")
				assert.Equal(t, item.expected, val, "should be equal")
			})
		}
	})

	t.Run("errors", func(t *testing.T) {
		table := []struct {
			expected string
			input    []string
		}{
			{
				"the dataset should not be empty",
				[]string{},
			},
			{
				"the dataset should not be empty",
				nil,
			},
			{
				"strconv.ParseFloat: parsing \"\": invalid syntax",
				[]string{"", "2,2", "19.0", "19.23"},
			},
			{
				"strconv.ParseFloat: parsing \"word\": invalid syntax",
				[]string{"word", "-50", "19.00", "-4.99"},
			},
			{
				"strconv.ParseFloat: parsing \"-5,0\": invalid syntax",
				[]string{"-5,0", "19.00", "-4.99"},
			},
		}
		for i, item := range table {
			t.Run(fmt.Sprintf("case #%d", i+1), func(t *testing.T) {
				val, err := stat.SortedDataSet(item.input)
				assert.Empty(t, val, "the output should be nil")
				assert.EqualError(t, err, item.expected, "should be error")
			})
		}
	})
}

func TestMedian(t *testing.T) {

	t.Run("success", func(t *testing.T) {
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
			{
				"-1.50",
				[]string{" -5.0", "-4.99", "2.0", "19.00"},
			},
		}
		for i, item := range table {
			t.Run(fmt.Sprintf("case #%d", i+1), func(t *testing.T) {
				val, err := stat.Median(item.input)
				assert.NoError(t, err, "should not be error")
				assert.Equal(t, item.expected, val, "should be equal")
			})
		}
	})

	t.Run("errors", func(t *testing.T) {
		table := []struct {
			expected string
			input    []string
		}{
			{
				"the dataset should not be empty",
				[]string{},
			},
			{
				"the dataset should not be empty",
				nil,
			},
			{
				"strconv.ParseFloat: parsing \"19,0\": invalid syntax",
				[]string{"19,0", "19.23", "24.0", "22.3"},
			},
			{
				"strconv.ParseFloat: parsing \"word\": invalid syntax",
				[]string{"word", "-4.99", "2.0", "19.00"},
			},
			{
				"strconv.ParseFloat: parsing \"-4,99\": invalid syntax",
				[]string{"-2", "-4,99", "2.0", "19.00"},
			},
		}
		for i, item := range table {
			t.Run(fmt.Sprintf("case #%d", i+1), func(t *testing.T) {
				val, err := stat.Median(item.input)
				assert.Empty(t, val, "the output should be nil")
				assert.EqualError(t, err, item.expected, "should be error")
			})
		}
	})

}
