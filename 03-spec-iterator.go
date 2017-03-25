package main

import (
	"io"
	"time"

	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/request"
	"github.com/tmornini/http-spec/response"
	"github.com/tmornini/http-spec/spec"
	"github.com/tmornini/http-spec/state"
)

func specIterator(state *state.State) {
	logger.Log("03-spec-iterator", state)

	for {
		state.Error = nil

		// desiredRequest, err := requestFromFile(state)
		desiredRequest, err := request.New(state)

		if err != nil {
			state.Error = err
			return
		}

		// expectedResponse, err := responseFromFile(state)
		expectedResponse, err := response.New(state)

		if err != nil && err != io.EOF {
			state.Error = err
			return
		}

		newSpec := spec.New()
		newSpec.DesiredRequest = desiredRequest
		newSpec.ExpectedResponse = expectedResponse

		state.Spec = newSpec

		// state.SpecTriplet = &specTriplet{
		// 	DesiredRequest:   desiredRequest,
		// 	ExpectedResponse: expectedResponse,
		// }

		desiredRequestSubstitor(state)

		newSpec.Duration = time.Since(newSpec.StartedAt)
		state.ResultsChannel <- *state
	}
}
