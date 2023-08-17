package lambda

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/janritter/aws-lambda-live-tuner/mocks/github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewLambda(t *testing.T) {
	testARN := "arn:aws:test"
	testArchitecture := "x86_64"
	testMemorySize := 256

	awsLambda := new(lambdaiface.MockLambdaAPI)
	awsLambda.On("GetFunctionConfiguration", mock.AnythingOfType("*lambda.GetFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(testArchitecture)},
		MemorySize:    aws.Int64(int64(testMemorySize)),
	}, nil)

	lambda, err := NewLambda(awsLambda, testARN, "")

	assert.Nil(t, err, "Expected no error")
	assert.NotEmpty(t, lambda, "Expected to be not empty")
	assert.NotNil(t, lambda, "Expected not to be nil")
	assert.Equal(t, testARN, lambda.Arn)
	assert.Equal(t, testMemorySize, lambda.PreTestMemory)
	assert.Equal(t, testArchitecture, lambda.Architecture)
}

func TestNewLambdaAlias(t *testing.T) {
	testARN := "arn:aws:test"
	testArchitecture := "x86_64"
	testMemorySize := 256

	awsLambda := new(lambdaiface.MockLambdaAPI)
	awsLambda.On("GetFunctionConfiguration", mock.AnythingOfType("*lambda.GetFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(testArchitecture)},
		MemorySize:    aws.Int64(int64(testMemorySize)),
	}, nil)

	awsLambda.On("GetAlias", mock.AnythingOfType("*lambda.GetAliasInput")).Return(&lambda.AliasConfiguration{
		FunctionVersion: aws.String("1"),
	}, nil)

	lambda, err := NewLambda(awsLambda, testARN, "live")

	assert.Nil(t, err, "Expected no error")
	assert.NotEmpty(t, lambda, "Expected to be not empty")
	assert.NotNil(t, lambda, "Expected not to be nil")
	assert.Equal(t, testARN, lambda.Arn)
	assert.Equal(t, testMemorySize, lambda.PreTestMemory)
	assert.Equal(t, testArchitecture, lambda.Architecture)
}

func TestNewLambdaError(t *testing.T) {
	testARN := "arn:aws:test"

	awsLambda := new(lambdaiface.MockLambdaAPI)
	awsLambda.On("GetFunctionConfiguration", mock.AnythingOfType("*lambda.GetFunctionConfigurationInput")).Return(nil, errors.New("Test error"))

	lambda, err := NewLambda(awsLambda, testARN, "")

	assert.Error(t, err, "Expected error")
	assert.Nil(t, lambda, "Expected to be nil")
}

func TestNewLambdaErrorAlias(t *testing.T) {
	testARN := "arn:aws:test"

	awsLambda := new(lambdaiface.MockLambdaAPI)
	awsLambda.On("GetFunctionConfiguration", mock.AnythingOfType("*lambda.GetFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{}, nil)
	awsLambda.On("GetAlias", mock.AnythingOfType("*lambda.GetAliasInput")).Return(nil, errors.New("Test error"))

	lambda, err := NewLambda(awsLambda, testARN, "live")

	assert.Error(t, err, "Expected error")
	assert.Nil(t, lambda, "Expected to be nil")
}
