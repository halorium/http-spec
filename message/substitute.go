package message

import "github.com/tmornini/http-spec/state"

func (message *Message) Substitute(state *state.State) {
	for _, singleLine := range message.allLines() {
		singleLine.Substitute(state)
	}
}
