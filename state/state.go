package state

import (
	"math/big"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/tmornini/http-spec/spec"
)

type State struct {
	ID           *big.Int
	LogFunctions bool
	LogState     bool
	// URLPrefix             string
	Pathnames              []string
	Pathname               string
	WaitGroup              *sync.WaitGroup
	ResultGathererChannel  chan State
	Tag                    string
	File                   *file
	Substitutions          map[string]string
	HTTPClient             *http.Client
	Spec                   *spec.Spec
	StartedAt              time.Time
	HTTPResponse           *http.Response
	Err                    error
	RegexpIdentifier       string
	SubstitutionIdentifier string
}

// func (context *context) log(tag string) {
// 	context.Tag = tag
//
// 	if context.LogFunctions {
// 		fmt.Println(tag)
// 	}
//
// 	if context.Logcontext {
// 		fmt.Printf("%#v\n", context)
// 	}
// }

func New() *State {
	s := &state{
		LogFunctions: false,
		Logcontext:   false,
		// URLPrefix:             prefix,
		Pathnames:              []string{},
		WaitGroup:              &sync.WaitGroup{},
		ResultGathererChannel:  make(chan state),
		StartedAt:              time.Now(),
		Substitutions:          map[string]string{},
		HTTPClient:             &http.Client{},
		RegexpIdentifier:       "⧆",
		SubstitutionIdentifier: "⧈",
	}

	for i, path := range os.Args {
		if i != 0 {
			s.Pathnames = append(s.Pathnames, path)
		}
	}

	return s
}
