package message

import (
	"github.com/tmornini/http-spec/line"
	"github.com/tmornini/http-spec/state"
)

type Header struct {
	Verb     string
	URI      string
	Protocol string
	Host     string
	Lines    []*line.Line
}

func NewHeader(state *state.State) *Header {
	return &Header{}
}

// func (message *Message) Header() string {
// 	headerLineTexts := []string{}
//
// 	for _, headerLine := range message.HeaderLines {
// 		headerLineTexts = append(headerLineTexts, headerLine.Text)
// 	}
//
// 	return strings.Join(headerLineTexts, "\n")
// }
