package message

import (
	"github.com/tmornini/http-spec/line"
	"github.com/tmornini/http-spec/state"
)

type Body struct {
	Lines []*line.Line
}

func NewBody(state *state.State) *Body {
	return &Body{}
}

// func (message *Message) Body() string {
// 	bodyLineTexts := []string{}
//
// 	for _, bodyLine := range message.BodyLines {
// 		bodyLineTexts = append(bodyLineTexts, bodyLine.Text)
// 	}
//
// 	return strings.Join(bodyLineTexts, "\n")
// }
