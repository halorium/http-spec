package message

import (
	"io"
	"strings"

	"github.com/tmornini/http-spec/line"
	"github.com/tmornini/http-spec/state"
)

func NewFromFile(state *state.State) (*Message, error) {
	message := &Message{}

	err := readFirstLine(message)

	if err != nil {
		return nil, err
	}

	err := readHostLine(message)

	if err != nil {
		return nil, err
	}

	err := readHeaderLines(message)

	if err != nil {
		return nil, err
	}

	err := readBodyLines(message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func readFirstLine(message *Message) error {
	firstLine, err := line.newFromFile(state)

	if err != nil {
		return err
	}

	lineArgs := strings.Split(firstLine.Text, " ")

	message.Header.Verb = lineArgs[0]
	message.Header.URI = lineArgs[1]
	message.Header.Protocol = lineArgs[2]
}

func readHostLine(message *Message) error {
	hostLine, err := line.newFromFile(state)

	if err != nil {
		return err
	}

	lineArgs := strings.Split(hostLine.Text, " ")

	message.Header.Host = lineArgs[1]
}

func readHeaderLines(message *Message) error {
	var headerLine *line.Line

	for {
		headerLine, err = line.newFromFile(state)

		if err != nil {
			return err
		}

		if headerLine.isEmpty() {
			message.BlankLine = headerLine

			break
		}

		message.Header.Lines = append(message.Header.Lines, headerLine)
	}
}

func readBodyLines(message *Message) error {
	var bodyLine *line

	for {
		bodyLine, err = line.newFromFile(state)

		if err == io.EOF || bodyLine.isBlank() {
			break
		}

		if err != nil {
			return err
		}

		message.Body.Lines = append(message.Body.Lines, bodyLine)
	}
}
