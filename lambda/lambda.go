package lambda

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/janritter/aws-lambda-live-tuner/helper"
	"golang.org/x/sync/semaphore"
)

// Prevent conflicts when changing memory on the function
var sem *semaphore.Weighted = semaphore.NewWeighted(int64(1))

func (l *Lambda) ChangeMemory(memory int) error {
	sem.Acquire(context.Background(), int64(1))

	retry := 0
	for retry <= 3 {
		_, err := l.awsLambda.UpdateFunctionConfiguration(&lambda.UpdateFunctionConfigurationInput{
			MemorySize:   aws.Int64(int64(memory)),
			FunctionName: aws.String(l.Arn),
		})

		if err == nil {
			sem.Release(int64(1))
			helper.LogNotice("Changed Lambda memory to: %d", memory)
			return nil
		}

		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lambda.ErrCodeResourceConflictException:
				helper.LogWarn("Lambda ResourceConflictException on memory update to %dMB - try: %d - Waiting 5 seconds", memory, retry+1)
				time.Sleep(time.Second * 5)
			default:
				sem.Release(int64(1))
				helper.LogError("Failed to change memory: %s", err)
				return err
			}
		}

		retry++
	}
	sem.Release(int64(1))

	helper.LogError("Maximum number of retries to change Lambda memory exceeded")
	return errors.New("maximum number of retries to change Lambda memory exceeded")
}

func (l *Lambda) ResetMemory() error {
	helper.LogNotice("Changing Lambda memory to pre-test value of %dMB", l.PreTestMemory)
	err := l.ChangeMemory(l.PreTestMemory)
	return err
}
