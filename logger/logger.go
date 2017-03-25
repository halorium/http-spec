package logger

import "fmt"

func Log(tag string, state interface{}) {
	logMessage := make(map[string]interface{})
	logMessage["tag"] = tag
	logMessage["state"] = state

	fmt.Printf("%#v\n", state)

	// state.Tag = tag
	//
	// if state.LogFunctions {
	// 	fmt.Println(tag)
	// }
	//
	// if state.LogState {
	// 	fmt.Printf("%#v\n", state)
	// }
}
