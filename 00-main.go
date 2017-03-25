package main

import (
	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/state"
)

const regexpIdentifier = "⧆"
const substitutionIdentifier = "⧈"

func main() {
	// var prefix string

	// startedAt := time.Now()

	// flag.StringVar(
	// 	&prefix,
	// 	"prefix",
	// 	"http://localhost:80",
	// 	"prefix for request URLs",
	// )
	//
	// flag.Parse()

	// filePaths := []string{}
	//
	// for i, path := range os.Args {
	// 	if i != 0 {
	// 		filePaths = append(filePaths, path)
	// 	}
	// }

	// filePaths := os.Args[1]

	// fmt.Println(filePaths)

	// fmt.Println(os.Args[1])

	// fmt.Println(flag.Args())

	thisState := state.New()

	// context := &context{
	// 	LogFunctions:          false,
	// 	Logcontext:              false,
	// 	URLPrefix:             prefix,
	// 	Pathnames:             flag.Args(),
	// 	WaitGroup:             &sync.WaitGroup{},
	// 	ResultsChannel: make(chan context),
	// 	StartedAt:             startedAt,
	// }

	// pathnames := []string{}

	// for i, path := range os.Args {
	// 	if i != 0 {
	// 		thisState.Pathnames = append(thisState.Pathnames, path)
	// 	}
	// }

	logger.Log("00-main", thisState)

	go resultGatherer(*thisState)

	specFileScatter(thisState)

	thisState.WaitGroup.Wait()

	close(thisState.ResultsChannel)

	select {}
}
