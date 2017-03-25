package line

import (
	"fmt"
	"strings"
)

func (line *Line) validate() error {
	if line.IsBlank() ||
		line.IsEmpty() ||
		line.IsRequest() ||
		line.IsResponse() ||
		line.IsComment() {
		return nil
	}

	return fmt.Errorf("malformed line: %s", line.Text)
}

func (line *Line) IsBlank() bool {
	return line.InputText == ""
}

func (line *Line) IsEmpty() bool {
	return line.InputText != "" && line.Text == ""
}

func (line *Line) IsRequest() bool {
	return line.IOPrefix != "" &&
		strings.HasPrefix(string(line.IOPrefix[0]), ">")
}

func (line *Line) IsResponse() bool {
	return line.IOPrefix != "" &&
		strings.HasPrefix(string(line.IOPrefix[0]), "<")
}

func (line *Line) IsComment() bool {
	return line.IOPrefix != "" &&
		strings.HasPrefix(string(line.IOPrefix[0]), "#")
}
