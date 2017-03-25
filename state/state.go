package state

import (
	"math/big"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/tmornini/http-spec/file"
)

var RegexpIdentifier = "⧆"
var SubstitutionIdentifier = "⧈"

type State struct {
	ID           *big.Int
	LogFunctions bool
	LogState     bool
	// URLPrefix             string
	Pathnames      []string
	Pathname       string
	WaitGroup      *sync.WaitGroup
	ResultsChannel chan State
	Tag            string
	File           *file.File
	Substitutions  map[string]string
	HTTPClient     *http.Client
	// Spec                   *spec.Spec
	Spec         interface{}
	StartedAt    time.Time
	HTTPResponse *http.Response
	Error        error
}

func New() *State {
	s := &State{
		LogFunctions: false,
		LogState:     false,
		// URLPrefix:             prefix,
		Pathnames:      []string{},
		WaitGroup:      &sync.WaitGroup{},
		ResultsChannel: make(chan State),
		StartedAt:      time.Now(),
		Substitutions:  map[string]string{},
		HTTPClient:     &http.Client{},
	}

	for i, path := range os.Args {
		if i != 0 {
			s.Pathnames = append(s.Pathnames, path)
		}
	}

	return s
}
