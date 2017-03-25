package main

import "strings"

type message struct {
	RequestLine   *line
	HeaderLines []*line
	BlankLine   *line
	BodyLines   []*line
}

// func messageFromFile(context *context) (*message, error) {
// 	RequestLine, err := newLineFromFile(context)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var headerLine *line
// 	var headerLines []*line
// 	var emptyLine *line
//
// 	for {
// 		headerLine, err = newLineFromFile(context)
//
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		if headerLine.isEmpty() {
// 			emptyLine = headerLine
//
// 			break
// 		}
//
// 		headerLines = append(headerLines, headerLine)
// 	}
//
// 	var bodyLine *line
// 	var bodyLines []*line
//
// 	for {
// 		bodyLine, err = newLineFromFile(context)
//
// 		if err == io.EOF || bodyLine.isBlank() {
// 			break
// 		}
//
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		bodyLines = append(bodyLines, bodyLine)
// 	}
//
// 	return &message{
// 		RequestLine,
// 		headerLines,
// 		emptyLine,
// 		bodyLines,
// 	}, nil
// }

func (message *message) allLines() []*line {
	var allLines []*line

	allLines = append(allLines, message.RequestLine)
	allLines = append(allLines, message.HeaderLines...)
	allLines = append(allLines, message.BodyLines...)

	return allLines
}

func (message *message) AllHeaderAndBodyLines() []*line {
	var AllHeaderAndBodyLines []*line

	AllHeaderAndBodyLines = append(AllHeaderAndBodyLines, message.HeaderLines...)
	AllHeaderAndBodyLines = append(AllHeaderAndBodyLines, message.BodyLines...)

	return AllHeaderAndBodyLines
}

func (message *message) substitute(context *context) {
	for _, line := range message.allLines() {
		line.substitute(context)
	}
}

func (message *message) Header() string {
	headerLineTexts := []string{}

	for _, headerLine := range message.HeaderLines {
		headerLineTexts = append(headerLineTexts, headerLine.Text)
	}

	return strings.Join(headerLineTexts, "\n")
}

func (message *message) Body() string {
	bodyLineTexts := []string{}

	for _, bodyLine := range message.BodyLines {
		bodyLineTexts = append(bodyLineTexts, bodyLine.Text)
	}

	return strings.Join(bodyLineTexts, "\n")
}
