package file

import "strings"

func (f *File) ReadLine() (string, error) {
	inputText, err := f.ReadString(byte('\n'))

	if err != nil {
		return "", err
	}

	f.LineNumber++

	inputText = strings.TrimSpace(inputText)

	return inputText, nil
}
