package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/spec"
	"github.com/tmornini/http-spec/state"
)

func desiredRequestSender(state *state.State) {
	logger.Log("07-desired-request-sender", state)

	thisSpec := state.Spec.(*spec.Spec)

	body := ioutil.NopCloser(strings.NewReader(thisSpec.DesiredRequest.Body() + "\n"))

	request, err := http.NewRequest(
		thisSpec.DesiredRequest.Method(),
		thisSpec.DesiredRequest.Host()+thisSpec.DesiredRequest.Path(),
		body,
	)

	if err != nil {
		state.Error = err
		return
	}

	for _, headerLine := range thisSpec.DesiredRequest.HeaderLines {
		parts := strings.Split(headerLine.Text, ":")

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		request.Header.Add(key, value)
	}

	thisSpec.StartedAt = time.Now()

	state.HTTPResponse, err = state.HTTPClient.Do(request)

	if err != nil {
		state.Error = err
		return
	}

	actualResponseReceiver(state)
}
