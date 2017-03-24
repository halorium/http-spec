package request

import (
	"github.com/tmornini/http-spec/message"
	"github.com/tmornini/http-spec/state"
)

type Request struct {
	*message
}

func New(state *state.State) (*Request, error) {
	message, err := message.New(state)

	if err != nil {
		return nil, err
	}

	return &Request{message}, nil
}
