package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"

	"github.com/tmornini/http-spec/line"
	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/message"
	"github.com/tmornini/http-spec/spec"
	"github.com/tmornini/http-spec/state"
)

func actualResponseReceiver(state *state.State) {
	logger.Log("08-actual-response-receiver", state)

	message := &message.Message{}

	responseLineNumber := 1

	version := state.HTTPResponse.Proto

	parts := strings.Split(state.HTTPResponse.Status, " ")

	statusCode := parts[0]
	reasonPhrase := strings.Join(parts[1:], " ")

	statusLineText := "< " + version + " " + statusCode + " " + reasonPhrase

	statusLine, err := line.NewFromText("response", responseLineNumber, statusLineText)

	if err != nil {
		state.Error = err
		return
	}

	responseLineNumber++

	message.RequestLine = statusLine

	var headerNames []string

	for headerName := range state.HTTPResponse.Header {
		headerNames = append(headerNames, headerName)
	}

	var headerLine *line.Line
	var headerLines []*line.Line

	sort.Strings(headerNames)

	for _, name := range headerNames {
		headerText := "< " + name + ": " + state.HTTPResponse.Header.Get(name)

		headerLine, err = line.NewFromText("response", responseLineNumber, headerText)

		if err != nil {
			state.Error = err
			return
		}

		headerLines = append(headerLines, headerLine)

		responseLineNumber++
	}

	message.HeaderLines = headerLines

	message.BlankLine, err = line.NewFromText("response", responseLineNumber, "<")

	if err != nil {
		state.Error = err
		return
	}

	responseLineNumber++

	scanner := bufio.NewScanner(state.HTTPResponse.Body)

	var bodyLine *line.Line
	var bodyLines []*line.Line

	for scanner.Scan() {
		if scanner.Err() != nil {
			state.Error = scanner.Err()
			return
		}

		bodyLine, err = line.NewFromText("response", responseLineNumber, "< "+scanner.Text())

		if err != nil {
			state.Error = err
			return
		}

		bodyLines = append(bodyLines, bodyLine)

		responseLineNumber++
	}

	state.HTTPResponse.Body.Close()

	message.BodyLines = bodyLines

	thisSpec := state.Spec.(*spec.Spec)

	thisSpec.ActualResponse.Message = message

	if thisSpec.IsRequestOnly() {
		state.Error = fmt.Errorf("no expected response")
		return
	}

	responseComparitor(state)
}
