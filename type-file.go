package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type file struct {
	*bufio.Reader
	PathName   string
	OSFile     *os.File
	LineNumber int
}

func (f *file) readLine() (string, error) {
	inputText, err := f.ReadString(byte('\n'))

	if err != nil {
		return "", err
	}

	f.LineNumber++

	inputText = strings.TrimSpace(inputText)

	return inputText, nil
}

func (f *file) String() string {
	return fmt.Sprintf("[%s:%3d]", f.PathName, f.LineNumber)
}

// func NewFile(context *context) (*file, error) {
// 	osFile, err := os.Open(context.Pathname)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &file{
// 		bufio.NewReader(osFile),
// 		context.Pathname,
// 		osFile,
// 		0,
// 	}, nil
// }
