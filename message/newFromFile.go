package message

import (
	"io"

	"github.com/tmornini/http-spec/line"
	"github.com/tmornini/http-spec/state"
)

func NewFromFile(state *state.State) (*Message, error) {
	message := &Message{}
	var err error

	err = readRequestLine(message, state)

	if err != nil {
		return nil, err
	}

	// err = readHostLine(message, state)
	//
	// if err != nil {
	// 	return nil, err
	// }

	err = readHeaderLines(message, state)

	if err != nil {
		return nil, err
	}

	err = readBodyLines(message, state)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func readRequestLine(message *Message, state *state.State) error {
	requestLine, err := line.NewFromFile(state)

	if err != nil {
		return err
	}

	message.RequestLine = requestLine

	return nil

	// lineArgs := strings.Split(requestLine.Text, " ")
	//
	// message.Header.Verb = lineArgs[0]
	// message.Header.URI = lineArgs[1]
	// message.Header.Protocol = lineArgs[2]
}

// func readHostLine(message *Message, state *state.State) error {
// 	hostLine, err := line.NewFromFile(state)
//
// 	if err != nil {
// 		return err
// 	}
//
//
// 	lineArgs := strings.Split(hostLine.Text, " ")
//
// 	message.Header.Host = lineArgs[1]
// }

func readHeaderLines(message *Message, state *state.State) error {
	var headerLine *line.Line
	var err error

	for {
		headerLine, err = line.NewFromFile(state)

		if err != nil {
			return err
		}

		if headerLine.IsEmpty() {
			message.BlankLine = headerLine

			break
		}

		message.HeaderLines = append(message.HeaderLines, headerLine)
	}

	return nil
}

func readBodyLines(message *Message, state *state.State) error {
	var bodyLine *line.Line
	var err error

	for {
		bodyLine, err = line.NewFromFile(state)

		if err == io.EOF || bodyLine.IsBlank() {
			break
		}

		// if err != nil {
		// 	return err
		// }

		message.BodyLines = append(message.BodyLines, bodyLine)
	}

	return err
}
