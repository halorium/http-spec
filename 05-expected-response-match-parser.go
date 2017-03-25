package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/matcher"
	"github.com/tmornini/http-spec/spec"
	"github.com/tmornini/http-spec/state"
)

func expectedResponseMatchParser(thisState *state.State) {
	logger.Log("05-expected-response-match-parser", thisState)

	thisSpec := thisState.Spec.(*spec.Spec)

	for _, line := range thisSpec.ExpectedResponse.AllHeaderAndBodyLines() {
		parts := strings.Split(line.Text, state.RegexpIdentifier)

		count := len(parts)

		if count == 1 {
			continue
		}

		if count == 0 || count == 2 || count == 3 || (count-4)%3 != 0 {
			// errorHandler(
			// 	state,
			// fmt.Errorf(
			// 	"regexp must be formed %soptional-capture-name%sregexp%s",
			// 	state.RegexpIdentifier,
			// 	state.RegexpIdentifier,
			// 	state.RegexpIdentifier,
			// ),
			// )
			thisState.Error = fmt.Errorf("regexp must be formed %soptional-capture-name%sregexp%s", state.RegexpIdentifier, state.RegexpIdentifier, state.RegexpIdentifier)
			return
		}

		var reg *regexp.Regexp
		var err error

		if parts[0] != "" {
			reg, err = regexp.Compile(regexp.QuoteMeta(parts[0]))

			if err != nil {
				thisState.Error = err
				return
			}

			line.RegexpNames = append(line.RegexpNames, ":prefix")
			line.Regexps = append(line.Regexps, reg)
		}

		matchers := matcher.New()

		for i := 1; i < count-1; i += 3 {
			reg = matchers[parts[i+1]]

			if reg == nil {
				reg, err = regexp.Compile("(" + parts[i+1] + ")")

				if err != nil {
					thisState.Error = err
					return
				}
			}

			// reString := "("
			//
			// switch parts[i+1] {
			// case ":date":
			// 	re = matchers[":date"]
			// case ":b62:22":
			// 	reString += "[0-9A-Za-z]{22}"
			// case ":uuid":
			// 	reString +=
			// 		"[[:xdigit:]]{8}-" +
			// 			"[[:xdigit:]]{4}-" +
			// 			"[[:xdigit:]]{4}-" +
			// 			"[[:xdigit:]]{4}-" +
			// 			"[[:xdigit:]]{12}"
			// default:
			// 	reString += parts[i+1]
			// }
			//
			// reString += ")"

			// re, err = regexp.Compile(reString)
			//
			// if errorHandler(state, err) {
			// 	return
			// }

			line.RegexpNames = append(line.RegexpNames, parts[i])
			line.Regexps = append(line.Regexps, reg)

			if parts[i+2] != "" {
				reg, err = regexp.Compile(regexp.QuoteMeta(parts[i+2]))

				if err != nil {
					thisState.Error = err
					return
				}

				line.RegexpNames = append(line.RegexpNames, ":postfix")
				line.Regexps = append(line.Regexps, reg)
			}
		}

		if err != nil {
			thisState.Error = err
			return
		}
	}

	expectedResponseSubstituter(thisState)
}
