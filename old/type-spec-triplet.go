package main

// import (
// 	"fmt"
// 	"strings"
// 	"time"
// )
//
// type specTriplet struct {
// 	DesiredRequest   *request
// 	ExpectedResponse *response
// 	ActualResponse   *response
// 	StartedAt        time.Time
// 	Duration         time.Duration
// }
//
// func (specTriplet *specTriplet) isRequestOnly() bool {
// 	return specTriplet.ExpectedResponse == nil
// }
//
// func (specTriplet *specTriplet) String() string {
// 	result := []string{}
//
// 	if specTriplet.DesiredRequest != nil {
// 		result =
// 			append(
// 				result,
// 				fmt.Sprintf(
// 					"%s:%d",
// 					specTriplet.DesiredRequest.RequestLine.PathName,
// 					specTriplet.DesiredRequest.RequestLine.LineNumber,
// 				),
// 			)
// 	}
//
// 	if specTriplet.ExpectedResponse != nil {
// 		result =
// 			append(
// 				result,
// 				fmt.Sprintf(
// 					"%d",
// 					specTriplet.ExpectedResponse.RequestLine.LineNumber,
// 				),
// 			)
// 	}
//
// 	return "[" + strings.Join(result, ":") + "]"
// }
