package main

import (
	"encoding/json"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/database"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/fn"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/repos"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/types"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func CreatePostHandler(req fn.Request) (fn.Response, error) {
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
	newPost, err := repo.CreatePost(post)
	if err != nil {
		return fn.MakeInternalServerErrorResponse(err.Error()), nil
	}
	return fn.MakeLambdaResponse(types.LambdaResponse{
		Error:   false,
		Message: "success",
		Data:    newPost,
	}.String()), nil
}

func main() {
	lambda.Start(CreatePostHandler)
}