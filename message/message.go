package message

import (
	"github.com/tmornini/http-spec/line"
	"github.com/tmornini/http-spec/state"
)

type Message struct {
	RequestLine *line.Line
	HeaderLines []*line.Line
	BlankLine   *line.Line
	BodyLines   []*line.Line
}

func New(state *state.State) (*Message, error) {
	return NewFromFile(state)
}
