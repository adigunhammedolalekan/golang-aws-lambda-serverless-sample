package main

import (
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/database"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/fn"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/repos"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/types"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"strconv"
)

func GetPostHandler(req fn.Request) (fn.Response, error) {
	evId := req.PathParameters["id"]
	id, err := strconv.Atoi(evId)
	if err != nil {
		return fn.MakeBadRequestResponse("bad request: event id is missing"), nil
	}

	db, err := database.New()
	if err != nil {
		return fn.MakeInternalServerErrorResponse("service unavailable"), nil
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println("failed to close database; ", err)
		}
	}()
	repo := repos.NewPostRepository(db.DB())
	p, err := repo.GetPost(uint(id))
	if err != nil {
		return fn.Make404RequestResponse(err.Error()), nil
	}
	return fn.MakeLambdaResponse(types.LambdaResponse{
		Error:   false,
		Message: "success",
		Data:    p,
	}.String()), nil
}

func main() {
	lambda.Start(GetPostHandler)
}
