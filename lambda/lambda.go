package lambda

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/janritter/aws-lambda-live-tuner/helper"
)

func (l *Lambda) ChangeMemory(memory int) error {
	_, err := l.awsLambda.UpdateFunctionConfiguration(&lambda.UpdateFunctionConfigurationInput{
		MemorySize:   aws.Int64(int64(memory)),
		FunctionName: aws.String(l.Arn),
	})
	if err != nil {
		helper.LogError("Failed to change memory: %s", err)
		return err
	}
	helper.LogNotice("Changed Lambda memory to: %d", memory)
	return nil
}

func (l *Lambda) ResetMemory() error {
	helper.LogNotice("Changing Lambda memory to pre-test value of %dMB", l.PreTestMemory)
	err := l.ChangeMemory(l.PreTestMemory)
	return err
}
