package message

import "github.com/tmornini/http-spec/state"

func (message *Message) substitute(state *state.State) {
	for _, line := range message.allLines() {
		line.substitute(state)
	}
}
