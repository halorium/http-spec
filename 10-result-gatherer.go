package main

import (
	"fmt"
	"math/big"
	"os"
	"time"
)

func resultGatherer(context context) {
	context.log("10 result-gatherer")

	success := true

	successCount := 0
	failureCount := 0

	outputs := map[*big.Int]string{}

	for completedcontext := range context.ResultGathererChannel {
		if completedcontext.Err == nil {
			// SUCCESS
			successCount++
			outputs[completedcontext.ID] +=
				fmt.Sprintf(
					"%s%s%s %s\n",
					Green,
					completedcontext.SpecTriplet.String(),
					Reset,
					completedcontext.SpecTriplet.Duration.String(),
				)
		} else {
			// FAILURE
			success = false
			failureCount++

			location := ""
			response := ""

			if completedcontext.File == nil {
				// file open failure
				location += "[" + completedcontext.Pathname + "]"
			} else {
				if completedcontext.SpecTriplet == nil {
					// request/response parsing failure
					location += completedcontext.File.String()
				} else {
					// request/response matching failure
					location +=
						completedcontext.SpecTriplet.String() + " " +
							completedcontext.SpecTriplet.Duration.String()

					if completedcontext.SpecTriplet.ActualResponse != nil {
						response = completedcontext.SpecTriplet.ActualResponse.String() + "\n"
					}
				}
			}

			outputs[completedcontext.ID] +=
				fmt.Sprintf(
					"%s%s%s %s\n%s\n",
					Red,
					location,
					Reset,
					completedcontext.Err.Error(),
					response,
				)
		}
	}

	duration := time.Since(context.StartedAt)

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
