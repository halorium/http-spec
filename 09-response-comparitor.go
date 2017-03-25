package main

import (
	"fmt"

	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/spec"
	"github.com/tmornini/http-spec/state"
)

func responseComparitor(state *state.State) {
	logger.Log("09-response-comparitor", state)

	thisSpec := state.Spec.(*spec.Spec)

	expectedResponseLines := thisSpec.ExpectedResponse.AllHeaderAndBodyLines()
	expectedCount := len(expectedResponseLines)

	actualResponseLines := thisSpec.ActualResponse.AllHeaderAndBodyLines()
	actualCount := len(actualResponseLines)

	if actualCount != expectedCount {
		state.Error = fmt.Errorf("response line count(%d) differs from expected line count (%d)", actualCount, expectedCount)
		return
	}

	for i, expectedResponseLine := range expectedResponseLines {
		err := expectedResponseLine.Compare(state, actualResponseLines[i])

		if err != nil {
			state.Error = err
			return
		}
	}
}
