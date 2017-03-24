package line

import (
	"fmt"
	"strings"
)

func (line *Line) validate() error {
	if line.isBlank() ||
		line.isEmpty() ||
		line.isRequest() ||
		line.isResponse() ||
		line.isComment() {
		return nil
	}

	return fmt.Errorf("malformed line: %s", line.String())
}

func (line *Line) isBlank() bool {
	return line.InputText == ""
}

func (line *Line) isEmpty() bool {
	return line.InputText != "" && line.Text == ""
}

func (line *Line) isRequest() bool {
	return line.IOPrefix != "" &&
		strings.HasPrefix(string(line.IOPrefix[0]), ">")
}

func (line *Line) isResponse() bool {
	return line.IOPrefix != "" &&
		strings.HasPrefix(string(line.IOPrefix[0]), "<")
}

func (line *Line) isComment() bool {
	return line.IOPrefix != "" &&
		strings.HasPrefix(string(line.IOPrefix[0]), "#")
}
