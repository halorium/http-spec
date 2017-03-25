package request

import (
	"strings"

	"github.com/tmornini/http-spec/message"
	"github.com/tmornini/http-spec/state"
)

type Request struct {
	*message.Message
}

func New(state *state.State) (*Request, error) {
	message, err := message.New(state)

	if err != nil {
		return nil, err
	}

	return &Request{message}, nil
}

func (request *Request) Method() string {
	return strings.Split(request.RequestLine.Text, " ")[0]
}

func (request *Request) Host() string {
	return strings.Split(request.HeaderLines[0].Text, " ")[1]
}

func (request *Request) Path() string {
	return strings.Split(request.RequestLine.Text, " ")[1]
}

func (request *Request) Version() string {
	return strings.Split(request.RequestLine.Text, " ")[2]
}

func (request *Request) String() string {
	lineStrings := []string{}

	lineStrings = append(lineStrings, request.RequestLine.String())

	for _, l := range request.HeaderLines {
		lineStrings = append(lineStrings, l.String())
	}

	lineStrings = append(lineStrings, request.BlankLine.String())

	if request.BodyLines != nil {
		for _, l := range request.BodyLines {
			lineStrings = append(lineStrings, l.String())
		}
	}

	return strings.Join(lineStrings, "\n")
}
