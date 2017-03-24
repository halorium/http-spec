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

func specTripletIterator(state *state.State) {
	logger.Log("03-spec-triplet-iterator", state)

	for {
		state.Err = nil

		// desiredRequest, err := requestFromFile(state)
		desiredRequest, err := request.New(state)

		if err != nil {
			state.Err = err
			state.ResultGathererChannel <- state
			return
		}

		// expectedResponse, err := responseFromFile(state)
		expectedResponse, err := response.New(state)

		if err != nil && err != io.EOF {
			state.Err = err
			state.ResultGathererChannel <- state
			return
		}

		state.Spec = spec.New()
		state.Spec.DesiredRequest = desiredRequest
		state.Spec.ExpectedResponse = expectedResponse

		// state.SpecTriplet = &specTriplet{
		// 	DesiredRequest:   desiredRequest,
		// 	ExpectedResponse: expectedResponse,
		// }

		desiredRequestSubstitor(state)

		state.Spec.Duration = time.Since(state.Spec.StartedAt)
		state.ResultGathererChannel <- *state
	}
}
