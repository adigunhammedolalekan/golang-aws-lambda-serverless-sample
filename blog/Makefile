.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/create create/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/delete delete/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/edit edit/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get get/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
