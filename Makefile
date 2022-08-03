prepare:
	go mod download

build: prepare
	go build -o ./bin/aws-lambda-live-tuner -v -ldflags "-X github.com/janritter/aws-lambda-live-tuner/cmd.gitSha=`git rev-parse HEAD` -X github.com/janritter/aws-lambda-live-tuner/cmd.buildTime=`date +'%Y-%m-%d_%T'` -X github.com/janritter/aws-lambda-live-tuner/cmd.version=LOCAL_BUILD"

tests:
	go test ./... -v  --cover

run:
	go run main.go
