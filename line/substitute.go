package line

import (
	"fmt"
	"strings"

	"github.com/tmornini/http-spec/state"
)

func (line *Line) Substitute(thisState *state.State) error {
	parts := strings.Split(line.Text, state.SubstitutionIdentifier)

	count := len(parts)

	if count == 1 {
		return nil
	}

	if count == 0 || count == 2 || (count-3)%2 != 0 {
		return fmt.Errorf("malformed substitution: %s", line.Text)
	}

	substitutedLine := parts[0]

	for i := 1; i < count-1; i += 2 {
		substitution, ok := thisState.Substitutions[parts[i]]

		if !ok {
			return fmt.Errorf("unknown tag: %v", parts[i])
		}

		substitutedLine += substitution + parts[i+1]
	}

	line.Text = substitutedLine

	return nil
}
