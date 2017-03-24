package logger

import (
	"fmt"

	"github.com/tmornini/http-spec/state"
)

func Log(tag string, state *state.State) {
	state.Tag = tag

	if state.LogFunctions {
		fmt.Println(tag)
	}

	if state.LogState {
		fmt.Printf("%#v\n", state)
	}
}
