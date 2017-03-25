package line

// state.File.PathName, state.File.LineNumber,
// pathName string, lineNumber int,
// line.NewFromText("response", responseLineNumber, statusLineText)
// newLineFromText(context.File.PathName, context.File.LineNumber, inputText)
func NewFromText(pathName string, lineNumber int, inputText string) (*Line, error) {
	ioPrefix, text := split(inputText)

	line := &Line{
		PathName:   pathName,
		LineNumber: lineNumber,
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
