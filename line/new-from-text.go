package line

import "github.com/tmornini/http-spec/state"

// state.File.PathName, state.File.LineNumber,
// pathName string, lineNumber int,
func newFromText(inputText string, state *state.State) (*Line, error) {
	ioPrefix, text := split(inputText)

	line := &Line{
		PathName:   state.File.PathName,
		LineNumber: state.File.LineNumber,
		InputText:  inputText,
		IOPrefix:   ioPrefix,
		Text:       text,
	}

	err := line.validate()

	if err != nil {
		return nil, err
	}

	return line, nil
}
