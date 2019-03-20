package formats

import (
	"encoding/csv"
	"io"
)

// CSV implements CSV structure
type CSV struct {
	Separator rune
}

// Parse implements CSV parser
func (c *CSV) Parse(r io.Reader) ([][]string, error) {
	csvFile := csv.NewReader(r)
	csvFile.Comma = c.Separator
	return csvFile.ReadAll()
}
