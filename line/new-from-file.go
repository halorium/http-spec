package line

import "github.com/tmornini/http-spec/state"

func newFromFile(state *state.State) (*Line, error) {
	for {
		inputText, err := state.File.ReadLine()

		if err != nil {
			return nil, err
		}

		line, err := newFromText(inputText, state)

		if err != nil {
			return nil, err
		}

		if !line.isComment() {
			break
		}
	}

	return line, nil
}
