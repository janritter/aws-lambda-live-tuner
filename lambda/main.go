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
	awsLambda      lambdaiface.LambdaAPI
	Arn            string
	Alias          string
	Architecture   string
	PreTestMemory  int
	PreTestVersion string
}

func NewLambda(awsLambda lambdaiface.LambdaAPI, arn string, alias string) (*Lambda, error) {
	result, err := awsLambda.GetFunctionConfiguration(&lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String(arn),
	})
	if err != nil {
		helper.LogError("Failed to get Lambda config: %s", err)
		return nil, err
	}

	if alias != "" {
		aliasResult, err := awsLambda.GetAlias(&lambda.GetAliasInput{
			FunctionName: aws.String(arn),
			Name:         aws.String(alias),
		})
		if err != nil {
			helper.LogError("Failed to get Lambda alias: %s", err)
			return nil, err
		}

		return &Lambda{
			awsLambda:      awsLambda,
			Arn:            arn,
			Alias:          alias,
			Architecture:   *result.Architectures[0],
			PreTestMemory:  int(*result.MemorySize),
			PreTestVersion: *aliasResult.FunctionVersion,
		}, nil
	}

	return &Lambda{
		awsLambda:      awsLambda,
		Arn:            arn,
		Alias:          alias,
		Architecture:   *result.Architectures[0],
		PreTestMemory:  int(*result.MemorySize),
		PreTestVersion: "",
	}, nil
}
