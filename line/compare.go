package line

import (
	"fmt"

	"github.com/tmornini/http-spec/state"
)

func (line *Line) Compare(state *state.State, otherLine *Line) error {
	if line.RegexpNames == nil && line.Text == otherLine.Text {
		return nil
	}

	if line.RegexpNames == nil && line.Text != otherLine.Text {
		return fmt.Errorf("%v != %v", line.Content(), otherLine.Content())
	}

	for i, regexpName := range line.RegexpNames {
		match := line.Regexps[i].FindString(otherLine.Text)

		if match == "" {
			return fmt.Errorf("%v !~ %v", line.Content(), otherLine.Content())
		}

		if regexpName != "" {
			state.Substitutions[regexpName] = match
		}
	}

	return nil
}
