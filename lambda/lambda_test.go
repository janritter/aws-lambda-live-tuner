package lambda

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/janritter/aws-lambda-live-tuner/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestChangeMemory(t *testing.T) {
	testARN := "arn:aws:test"
	testArchitecture := "x86_64"
	testMemorySize := 256

	awsLambda := new(mocks.LambdaAPI)
	awsLambda.On("GetFunctionConfiguration", mock.AnythingOfType("*lambda.GetFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(testArchitecture)},
		MemorySize:    aws.Int64(int64(testMemorySize)),
	}, nil)

	awsLambda.On("UpdateFunctionConfiguration", mock.AnythingOfType("*lambda.UpdateFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(testArchitecture)},
		MemorySize:    aws.Int64(int64(testMemorySize)),
	}, nil)

	mockLambda, err := NewLambda(awsLambda, testARN, "")

	assert.Nil(t, err, "Expected no error")
	assert.NotEmpty(t, mockLambda, "Expected to be not empty")

	err = mockLambda.ChangeMemory(512)
	assert.Nil(t, err, "Expected no error")

	awsLambda.AssertCalled(t, "UpdateFunctionConfiguration", &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String(testARN),
		MemorySize:   aws.Int64(512),
	})
}

func TestChangeMemoryError(t *testing.T) {
	testARN := "arn:aws:test"
	testArchitecture := "x86_64"
	testMemorySize := 256

	awsLambda := new(mocks.LambdaAPI)
	awsLambda.On("GetFunctionConfiguration", mock.AnythingOfType("*lambda.GetFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(testArchitecture)},
		MemorySize:    aws.Int64(int64(testMemorySize)),
	}, nil)

	awsLambda.On("UpdateFunctionConfiguration", mock.AnythingOfType("*lambda.UpdateFunctionConfigurationInput")).Return(nil, errors.New("test"))

	mockLambda, err := NewLambda(awsLambda, testARN, "")

	assert.Nil(t, err, "Expected no error")
	assert.NotEmpty(t, mockLambda, "Expected to be not empty")

	err = mockLambda.ChangeMemory(512)
	assert.NotNil(t, err, "Expected error")
	assert.Equal(t, "test", err.Error())
}

func TestChangeMemoryResourceConflictError(t *testing.T) {
	testARN := "arn:aws:test"
	testArchitecture := "x86_64"
	testMemorySize := 256

	awsLambda := new(mocks.LambdaAPI)
	awsLambda.On("GetFunctionConfiguration", mock.AnythingOfType("*lambda.GetFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(testArchitecture)},
		MemorySize:    aws.Int64(int64(testMemorySize)),
	}, nil)

	awsLambda.On("UpdateFunctionConfiguration", mock.AnythingOfType("*lambda.UpdateFunctionConfigurationInput")).Return(nil, awserr.New(lambda.ErrCodeResourceConflictException, "ResourceConflictException message", nil))

	mockLambda, err := NewLambda(awsLambda, testARN, "")
	assert.Nil(t, err, "Expected no error")
	assert.NotEmpty(t, mockLambda, "Expected to be not empty")

	err = mockLambda.ChangeMemory(512)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, err, errors.New("maximum number of retries to change Lambda memory exceeded"))
}

func TestChangeMemoryAlias(t *testing.T) {
	testARN := "arn:aws:test"
	testArchitecture := "x86_64"
	testMemorySize := 256
	newVersion := "1337"

	awsLambda := new(mocks.LambdaAPI)
	awsLambda.On("GetFunctionConfiguration", mock.AnythingOfType("*lambda.GetFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(testArchitecture)},
		MemorySize:    aws.Int64(int64(testMemorySize)),
	}, nil)

	awsLambda.On("UpdateFunctionConfiguration", mock.AnythingOfType("*lambda.UpdateFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(testArchitecture)},
		MemorySize:    aws.Int64(int64(testMemorySize)),
	}, nil)

	awsLambda.On("GetAlias", mock.AnythingOfType("*lambda.GetAliasInput")).Return(&lambda.AliasConfiguration{
		FunctionVersion: aws.String("1"),
	}, nil)

	awsLambda.On("PublishVersion", mock.AnythingOfType("*lambda.PublishVersionInput")).Return(&lambda.FunctionConfiguration{
		Version: aws.String(newVersion),
	}, nil)

	awsLambda.On("UpdateAlias", mock.AnythingOfType("*lambda.UpdateAliasInput")).Return(&lambda.AliasConfiguration{
		FunctionVersion: aws.String(newVersion),
	}, nil)

	mockLambda, err := NewLambda(awsLambda, testARN, "test")

	assert.Nil(t, err, "Expected no error")
	assert.NotEmpty(t, mockLambda, "Expected to be not empty")

	err = mockLambda.ChangeMemory(512)
	assert.Nil(t, err, "Expected no error")

	awsLambda.AssertCalled(t, "UpdateFunctionConfiguration", &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String(testARN),
		MemorySize:   aws.Int64(512),
	})

	awsLambda.AssertCalled(t, "UpdateAlias", &lambda.UpdateAliasInput{
		FunctionName:    aws.String(testARN),
		FunctionVersion: aws.String(newVersion),
		Name:            aws.String("test"),
	})
}
