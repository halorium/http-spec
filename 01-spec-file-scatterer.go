package main

import (
	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/state"
)

func specFileScatter(state *state.State) {
	logger.Log("01-spec-file-scatter", state)

	for _, pathname := range state.Pathnames {
		state.Pathname = pathname

		state.WaitGroup.Add(1)
		go specFileProcessor(*state)
	}
}
