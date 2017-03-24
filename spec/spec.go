package spec

import (
	"fmt"
	"strings"
	"time"

	"github.com/tmornini/http-spec/request"
	"github.com/tmornini/http-spec/response"
)

type Spec struct {
	DesiredRequest   *request.Request
	ExpectedResponse *response.Response
	ActualResponse   *response.Response
	StartedAt        time.Time
	Duration         time.Duration
}

func New() *Spec {
	return &Spec{}
}

func (spec *Spec) isRequestOnly() bool {
	return spec.ExpectedResponse == nil
}

func (spec *Spec) String() string {
	result := []string{}

	if spec.DesiredRequest != nil {
		result =
			append(
				result,
				fmt.Sprintf(
					"%s:%d",
					spec.DesiredRequest.FirstLine.PathName,
					spec.DesiredRequest.FirstLine.LineNumber,
				),
			)
	}

	if spec.ExpectedResponse != nil {
		result =
			append(
				result,
				fmt.Sprintf(
					"%d",
					spec.ExpectedResponse.FirstLine.LineNumber,
				),
			)
	}

	return "[" + strings.Join(result, ":") + "]"
}
