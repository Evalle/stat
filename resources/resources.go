package resources

import (
	"io"
	"os"
)

// TODO:
type LocalFile struct {
	FileName string
}

// Fetch implements fetching the file from the local filesystem
func (l *LocalFile) Fetch() (io.ReadCloser, error) {
	return os.Open(l.FileName)
}
