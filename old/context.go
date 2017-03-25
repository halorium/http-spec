package main

import (
	"fmt"
	"math/big"
	"net/http"
	"sync"
	"time"
)

type context struct {
	ID           *big.Int
	LogFunctions bool
	Logcontext   bool
	// URLPrefix             string
	Pathnames             []string
	Pathname              string
	WaitGroup             *sync.WaitGroup
	ResultsChannel chan context
	Tag                   string
	File                  *file
	Substitutions         map[string]string
	HTTPClient            *http.Client
	SpecTriplet           *specTriplet
	StartedAt             time.Time
	HTTPResponse          *http.Response
	Err                   error
}

func (context *context) log(tag string) {
	context.Tag = tag

	if context.LogFunctions {
		fmt.Println(tag)
	}

	if context.Logcontext {
		fmt.Printf("%#v\n", context)
	}
}

func newContext() *context {
	return &context{
		LogFunctions: false,
		Logcontext:   false,
		// URLPrefix:             prefix,
		Pathnames:             []string{},
		WaitGroup:             &sync.WaitGroup{},
		ResultsChannel: make(chan context),
		StartedAt:             time.Now(),
		Substitutions:         map[string]string{},
		HTTPClient:            &http.Client{},
	}
}
