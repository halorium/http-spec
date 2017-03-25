package response

import (
	"strings"

	"github.com/tmornini/http-spec/state"
)

func (response *Response) String() string {
	lineStrings := []string{}

	lineStrings = append(lineStrings, response.RequestLine.Content())

	for _, l := range response.HeaderLines {
		content := l.Content()

		if content[0:8] == "< Date: " {
			content =
				content[0:8] +
					state.RegexpIdentifier +
					state.RegexpIdentifier +
					":date" +
					state.RegexpIdentifier
		}

		lineStrings = append(lineStrings, content)
	}

	lineStrings = append(lineStrings, response.BlankLine.Content())

	if response.BodyLines != nil {
		for _, l := range response.BodyLines {
			lineStrings = append(lineStrings, l.Content())
		}
	}

	return strings.Join(lineStrings, "\n")
}
