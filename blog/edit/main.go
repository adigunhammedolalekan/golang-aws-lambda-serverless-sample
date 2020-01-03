package main

import (
	"encoding/json"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/database"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/fn"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/repos"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/types"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"strconv"
)

func EditPostHandler(req fn.Request) (fn.Response, error) {
	evId := req.PathParameters["id"]
	id, err := strconv.Atoi(evId)
	if err != nil {
		return fn.MakeBadRequestResponse("bad request: event id is missing"), nil
	}
	type payload struct {
		Title string `json:"title"`
		Body string `json:"body"`
		User string `json:"user"`
	}
	p := &payload{}
	if err := json.Unmarshal(req.BodyBytes(), p); err != nil {
		return fn.MakeBadRequestResponse(err.Error()), nil
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
	post := types.NewPost(p.Title, p.Body, p.User)
	updatedPost, err := repo.EditPost(uint(id), post)
	if err != nil {
		return fn.MakeInternalServerErrorResponse(err.Error()), nil
	}
	return fn.MakeLambdaResponse(types.LambdaResponse{
		Error:   false,
		Message: "post updated",
		Data:    updatedPost,
	}.String()), nil
}

func main() {
	lambda.Start(EditPostHandler)
}
