package changer

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"go.uber.org/zap"
)

func (c *Changer) ChangeMemory(lambdaARN string, memory int) error {
	_, err := c.lambda.UpdateFunctionConfiguration(&lambda.UpdateFunctionConfigurationInput{
		MemorySize:   aws.Int64(int64(memory)),
		FunctionName: aws.String(lambdaARN),
	})
	if err != nil {
		c.logger.Error("Failed to change memory: ", zap.Error(err))
	}
	c.logger.Infof("Changed Lambda memory to: %d", memory)
	return nil
}

func (c *Changer) GetCurrentMemoryValue(lambdaARN string) (int, error) {
	result, err := c.lambda.GetFunctionConfiguration(&lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String(lambdaARN),
	})
	if err != nil {
		c.logger.Error("Failed to get current memory value: ", zap.Error(err))
		return -1, err
	}

	return int(*result.MemorySize), nil
}
