package line

import "github.com/tmornini/http-spec/state"

func NewFromFile(state *state.State) (*Line, error) {
	var line *Line

	for {
		inputText, err := state.File.ReadLine()

		if err != nil {
			return nil, err
		}

		line, err = NewFromText(state.File.PathName, state.File.LineNumber, inputText)

		if err != nil {
			return nil, err
		}

		if !line.IsComment() {
			break
		}
	}

	return line, nil
}
