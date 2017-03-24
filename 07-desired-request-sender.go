package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/state"
)

func desiredRequestSender(state *state.State) {
	logger.Log("07-desired-request-sender", state)

	body := ioutil.NopCloser(strings.NewReader(state.Spec.DesiredRequest.Body() + "\n"))

	request, err := http.NewRequest(
		state.Spec.DesiredRequest.Method(),
		state.URLPrefix+state.Spec.DesiredRequest.Path(),
		body,
	)

	if errorHandler(state, err) {
		return
	}

	for _, headerLine := range state.Spec.DesiredRequest.HeaderLines {
		parts := strings.Split(headerLine.Text, ":")

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		request.Header.Add(key, value)
	}

	state.Spec.StartedAt = time.Now()

	state.HTTPResponse, err = state.HTTPClient.Do(request)

	if errorHandler(state, err) {
		return
	}

	actualResponseReceiver(state)
}
