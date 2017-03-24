package message

import (
	"strings"

	"github.com/tmornini/http-spec/line"
)

func (message *Message) allLines() []*line.Line {
	var lines []*line.Line

	lines = append(lines, message.FirstLine)
	lines = append(lines, message.HeaderLines...)
	lines = append(lines, message.BodyLines...)

	return lines
}

func (message *Message) allHeaderAndBodyLines() []*line.Line {
	var lines []*line.Line

	lines = append(lines, message.HeaderLines...)
	lines = append(lines, message.BodyLines...)

	return lines
}

func (message *Message) Header() string {
	headerLineTexts := []string{}

	for _, headerLine := range message.HeaderLines {
		headerLineTexts = append(headerLineTexts, headerLine.Text)
	}

	return strings.Join(headerLineTexts, "\n")
}

func (message *Message) Body() string {
	bodyLineTexts := []string{}

	for _, bodyLine := range message.BodyLines {
		bodyLineTexts = append(bodyLineTexts, bodyLine.Text)
	}

	return strings.Join(bodyLineTexts, "\n")
}
