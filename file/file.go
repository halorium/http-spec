package file

import (
	"bufio"
	"os"
)

type File struct {
	*bufio.Reader
	PathName   string
	OSFile     *os.File
	LineNumber int
}

func New(pathname string) (*File, error) {
	osFile, err := os.Open(pathname)

	if err != nil {
		return nil, err
	}

	return &File{
		bufio.NewReader(osFile),
		pathname,
		osFile,
		0,
	}, nil
}
