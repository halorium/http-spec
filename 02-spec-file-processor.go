package main

import (
	"crypto/rand"
	"math/big"

	"github.com/tmornini/http-spec/file"
	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/state"
)

func specFileProcessor(state state.State) {
	logger.Log("02-spec-file-processor", state)

	defer state.WaitGroup.Done()

	// osFile, err := os.Open(state.Pathname)
	//
	// if errorHandler(&state, err) {
	// 	state.ResultsChannel <- state
	//
	// 	return
	// }

	space := new(big.Int).Exp(big.NewInt(62), big.NewInt(22), nil)
	uuid, err := rand.Int(rand.Reader, space)

	if err != nil {
		panic(err)
	}

	state.ID = uuid

	// state.File = &file{
	// 	bufio.NewReader(osFile),
	// 	state.Pathname,
	// 	osFile,
	// 	0,
	// }

	file, err := file.New(state.Pathname)

	if err != nil {
		state.Error = err
		state.ResultsChannel <- state
		return
	}

	defer file.OSFile.Close()

	state.File = file

	// state.Substitutions = map[string]string{}
	//
	// state.HTTPClient = &http.Client{}

	specIterator(&state)
}
