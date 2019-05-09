// [x] rename reddata package to `data`
// [x] rename data.ReadData to data.Read
// [x] new package `resources`, implements Fetch: os.Open (struct)
// [x] new package `formats`, implements Parse: csv.NewReader (struct)
// [ ] in main, new struct that embeds structs w/ Fetch and Parse: injected into data.Read

package data

import (
	"io"
)

// FetchParser is an interface for Readdata ...TBA
type FetchParser interface {
	Fetch() (io.ReadCloser, error)
	Parse(io.Reader) ([][]string, error)
}

// Read reads the csv data from the file and returns rows
func Read(data FetchParser) ([][]string, error) {
	f, err := data.Fetch()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r, err := data.Parse(f)

	return r, err
}
