package line

import (
	"regexp"

	"github.com/tmornini/http-spec/state"
)

type Line struct {
	LineNumber  int
	PathName    string
	InputText   string
	IOPrefix    string
	Text        string
	RegexpNames []string
	Regexps     []*regexp.Regexp
}

func New(state *state.State) (*Line, error) {
	return newFromFile(state)
}
