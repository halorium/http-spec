package request

import (
	"github.com/tmornini/http-spec/message"
	"github.com/tmornini/http-spec/state"
)

func NewFromFile(state *state.State) (*Request, error) {
	message, err := message.NewFromFile(state)

	if err != nil {
		return nil, err
	}

	return &Request{message}, nil
}
