package main

import (
	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/state"
)

func desiredRequestSubstitor(state *state.State) {
	logger.Log("04-desired-request-substituter")

	state.Spec.DesiredRequest.substitute(state)

	if state.Spec.isRequestOnly() {
		desiredRequestSender(state)

		return
	}

	expectedResponseMatchParser(state)
}
