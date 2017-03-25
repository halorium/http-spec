package main

import (
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/tmornini/http-spec/logger"
	"github.com/tmornini/http-spec/spec"
	"github.com/tmornini/http-spec/state"
)

func resultGatherer(thisState state.State) {
	logger.Log("10-result-gatherer", thisState)

	success := true

	successCount := 0
	failureCount := 0

	outputs := map[*big.Int]string{}

	for completedState := range thisState.ResultsChannel {
		if completedState.Error == nil {
			// SUCCESS
			thisSpec := completedState.Spec.(*spec.Spec)
			successCount++
			outputs[completedState.ID] +=
				fmt.Sprintf(
					"%s%s%s %s\n",
					Green,
					thisSpec.String(),
					Reset,
					thisSpec.Duration.String(),
				)
		} else {
			// FAILURE
			success = false
			failureCount++

			location := ""
			response := ""

			if completedState.File == nil {
				// file open failure
				location += "[" + completedState.Pathname + "]"
			} else {
				if completedState.Spec == nil {
					// request/response parsing failure
					location += completedState.File.String()
				} else {
					// request/response matching failure
					thisSpec := completedState.Spec.(*spec.Spec)
					location +=
						thisSpec.String() + " " +
							thisSpec.Duration.String()

					if thisSpec.ActualResponse != nil {
						response = thisSpec.ActualResponse.String() + "\n"
					}
				}
			}

			outputs[completedState.ID] +=
				fmt.Sprintf(
					"%s%s%s %s\n%s\n",
					Red,
					location,
					Reset,
					completedState.Error.Error(),
					response,
				)
		}
	}

	duration := time.Since(thisState.StartedAt)

	fmt.Println()

	for _, result := range outputs {
		fmt.Print(result)
	}

	if !success {
		fmt.Printf(
			"%sFAILURE:%s %s+%d%s %s-%d%s %s\n",
			Red,
			Reset,
			Green,
			successCount,
			Reset,
			Red,
			failureCount,
			Reset,
			duration.String(),
		)

		os.Exit(1)
	}

	fmt.Printf(
		"\n%sSUCCESS:%s %s+%d%s %s\n",
		Green,
		Reset,
		Green,
		successCount,
		Reset,
		duration.String(),
	)

	os.Exit(0)
}
