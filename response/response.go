package response

import (
	"github.com/tmornini/http-spec/message"
	"github.com/tmornini/http-spec/state"
)

type Response struct {
	*message.Message
}

func New(state *state.State) (*Response, error) {
	message, err := message.New(state)

	if err != nil {
		return nil, err
	}

	return &Response{message}, nil
}
