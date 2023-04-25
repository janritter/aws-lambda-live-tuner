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

// Prevent conflicts when changing options on the function
var sem *semaphore.Weighted = semaphore.NewWeighted(int64(1))

func (l *Lambda) changeUnpublishedMemory(memory int) error {
	sem.Acquire(context.Background(), int64(1))

	retry := 0
	for retry <= 5 {
		retry++

		_, err := l.awsLambda.UpdateFunctionConfiguration(&lambda.UpdateFunctionConfigurationInput{
			MemorySize:   aws.Int64(int64(memory)),
			FunctionName: aws.String(l.Arn),
		})

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				if aerr.Code() == lambda.ErrCodeResourceConflictException {
					helper.LogWarn("Lambda ResourceConflictException on memory update to %dMB - try: %d - Waiting 2 seconds", memory, retry)
					time.Sleep(time.Second * 2)
					continue
				}
			}
			helper.LogError("Failed to change memory to %sMB", err)

			sem.Release(int64(1))
			return err
		} else {
			helper.LogNotice("Changed Lambda memory to %dMB", memory)

			sem.Release(int64(1))
			return nil
		}
	}
	helper.LogError("Maximum number of retries to change Lambda memory exceeded")

	sem.Release(int64(1))
	return errors.New("maximum number of retries to change Lambda memory exceeded")
}

func (l *Lambda) updateAlias(version string) error {
	sem.Acquire(context.Background(), int64(1))

	retry := 0
	for retry <= 5 {
		retry++

		_, err := l.awsLambda.UpdateAlias(&lambda.UpdateAliasInput{
			FunctionName:    aws.String(l.Arn),
			FunctionVersion: aws.String(version),
			Name:            aws.String(l.Alias),
		})

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				if aerr.Code() == lambda.ErrCodeResourceConflictException {
					helper.LogWarn("Lambda ResourceConflictException on alias %s update to version %s - try: %d - Waiting 2 seconds", l.Alias, version, retry)
					time.Sleep(time.Second * 2)
					continue
				}
			}
			helper.LogError("Failed to change alias %s to version %s", l.Alias, err)

			sem.Release(int64(1))
			return err
		} else {
			helper.LogNotice("Changed Lambda alias %s to version %s", l.Alias, version)

			sem.Release(int64(1))
			return nil
		}
	}
	helper.LogError("Maximum number of retries to change Lambda alias exceeded")

	sem.Release(int64(1))
	return errors.New("maximum number of retries to change Lambda alias exceeded")
}

func (l *Lambda) publishVersion() (string, error) {
	sem.Acquire(context.Background(), int64(1))

	retry := 0
	for retry <= 5 {
		retry++

		version, err := l.awsLambda.PublishVersion(&lambda.PublishVersionInput{
			FunctionName: aws.String(l.Arn),
		})

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				if aerr.Code() == lambda.ErrCodeResourceConflictException {
					helper.LogWarn("Lambda ResourceConflictException on publish new version - try: %d - Waiting 2 seconds", retry)
					time.Sleep(time.Second * 2)
					continue
				}
			}
			helper.LogError("Failed to publish new version %s", err)

			sem.Release(int64(1))
			return "", err
		} else {
			helper.LogNotice("Published new version %s of Lambda", *version.Version)

			sem.Release(int64(1))
			return *version.Version, nil
		}
	}
	helper.LogError("Maximum number of retries to publish new version exceeded")

	sem.Release(int64(1))
	return "", errors.New("maximum number of retries to publish new version exceeded")
}

func (l *Lambda) ChangeMemory(memory int) error {
	err := l.changeUnpublishedMemory(memory)
	if err != nil {
		return err
	}

	if l.Alias != "" {
		time.Sleep(time.Second * 2)
		version, err := l.publishVersion()
		if err != nil {
			return err
		}

		time.Sleep(time.Second * 2)
		err = l.updateAlias(version)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lambda) Reset() error {
	if l.Alias != "" {
		helper.LogNotice("Changing Lambda alias %s to pre-test version %s", l.Alias, l.PreTestVersion)

		err := l.updateAlias(l.PreTestVersion)
		if err != nil {
			return err
		}
	}

	helper.LogNotice("Changing Lambda memory to pre-test value of %dMB", l.PreTestMemory)
	err := l.changeUnpublishedMemory(l.PreTestMemory)
	return err
}
