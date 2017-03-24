package response

import "github.com/tmornini/http-spec/state"

type Response struct {
	*message
}

func New(state *state.State) (*Response, error) {
	message, err := message.New(state)

	if err != nil {
		return nil, err
	}

	return &Response{message}, nil
}
