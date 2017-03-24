package message

import (
	"github.com/tmornini/http-spec/line"
	"github.com/tmornini/http-spec/state"
)

type Message struct {
	Header    *Header
	BlankLine *line.Line
	Body      *Body
}

func New(state *state.State) (*Message, error) {
	return NewFromFile(state)
}
