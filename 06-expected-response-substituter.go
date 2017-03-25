package main

import (
	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/spec"
	"github.com/tmornini/http-spec/state"
)

func expectedResponseSubstituter(state *state.State) {
	logger.Log("06-expected-response-substituter", state)

	thisSpec := state.Spec.(*spec.Spec)

	thisSpec.ExpectedResponse.Substitute(state)

	desiredRequestSender(state)
}
