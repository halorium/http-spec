package main

import (
	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/spec"
	"github.com/tmornini/http-spec/state"
)

func desiredRequestSubstitor(state *state.State) {
	logger.Log("04-desired-request-substituter", state)

	thisSpec := state.Spec.(spec.Spec)

	thisSpec.DesiredRequest.Substitute(state)

	if thisSpec.IsRequestOnly() {
		desiredRequestSender(state)

		return
	}

	expectedResponseMatchParser(state)
}
