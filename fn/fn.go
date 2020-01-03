package fn

import (
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/types"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func (r Request) BodyBytes() []byte {
	return []byte(r.Body)
}

func MakeBadRequestResponse(message string) Response {
	r := types.LambdaResponse{
		Error:   true,
		Message: message,
	}
	return Response{
		StatusCode: http.StatusBadRequest,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: r.String(),
	}
}

func Make404RequestResponse(message string) Response {
	return Response{
		StatusCode: http.StatusNotFound,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: types.LambdaResponse{
			Error:   false,
			Message: message,
		}.String(),
	}
}

func MakeLambdaResponse(response string) Response {
	return Response{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: response,
	}
}

func MakeInternalServerErrorResponse(message string) Response {
	r := types.LambdaResponse{
		Error:   true,
		Message: message,
	}
	return Response{
		StatusCode: http.StatusInternalServerError,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: r.String(),
	}
}

func Make401Response(message string) Response {
	return Response{
		StatusCode: http.StatusUnauthorized,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       types.LambdaResponse{Error: false, Message: message}.String(),
	}
}
