package line

import (
	"fmt"
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
	return NewFromFile(state)
}

func (line *Line) String() string {
	return fmt.Sprintf("%s %s", line.Location(), line.Content())
}

func (line *Line) Location() string {
	return fmt.Sprintf("[%s:%3d]", line.PathName, line.LineNumber)
}

func (line *Line) Content() string {
	return fmt.Sprintf("%s %s", line.IOPrefix, line.Text)
}
