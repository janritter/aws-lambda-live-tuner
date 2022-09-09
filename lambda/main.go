package lambda

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/janritter/aws-lambda-live-tuner/helper"
)

type LambdaAPI interface {
	ChangeMemory(memory int) error
}

type Lambda struct {
	awsLambda     lambdaiface.LambdaAPI
	Arn           string
	Architecture  string
	PreTestMemory int
}

func NewLambda(awsLambda lambdaiface.LambdaAPI, arn string) (*Lambda, error) {
	result, err := awsLambda.GetFunctionConfiguration(&lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String(arn),
	})
	if err != nil {
		helper.LogError("Failed to get Lambda config: %s", err)
		return nil, err
	}

	return &Lambda{
		awsLambda:     awsLambda,
		Arn:           arn,
		Architecture:  *result.Architectures[0],
		PreTestMemory: int(*result.MemorySize),
	}, nil
}
