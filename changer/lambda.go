package changer

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/janritter/aws-lambda-live-tuner/helper"
)

func (c *Changer) ChangeMemory(lambdaARN string, memory int) error {
	_, err := c.lambda.UpdateFunctionConfiguration(&lambda.UpdateFunctionConfigurationInput{
		MemorySize:   aws.Int64(int64(memory)),
		FunctionName: aws.String(lambdaARN),
	})
	if err != nil {
		helper.LogError("Failed to change memory: %s", err)
		return err
	}
	helper.LogNotice("Changed Lambda memory to: %d", memory)
	return nil
}

func (c *Changer) GetCurrentMemoryValue(lambdaARN string) (int, error) {
	result, err := c.lambda.GetFunctionConfiguration(&lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String(lambdaARN),
	})
	if err != nil {
		helper.LogError("Failed to get current memory value: %s", err)
		return -1, err
	}

	return int(*result.MemorySize), nil
}

func (c *Changer) GetArchitecture(lambdaARN string) (string, error) {
	result, err := c.lambda.GetFunctionConfiguration(&lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String(lambdaARN),
	})
	if err != nil {
		helper.LogError("Failed to get architecture: %s", err)
		return "", err
	}

	return *result.Architectures[0], nil
}
