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

// Returns a mocked Lambda struct and the mocked AWS Lambda API
func getMockLambda() (*Lambda, *mocks.LambdaAPI) {
	testARN := "arn:aws:test"
	testArchitecture := "x86_64"
	testMemorySize := 256

	awsLambda := new(mocks.LambdaAPI)
	awsLambda.On("GetFunctionConfiguration", mock.AnythingOfType("*lambda.GetFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(testArchitecture)},
		MemorySize:    aws.Int64(int64(testMemorySize)),
	}, nil)

	mockLambda, err := NewLambda(awsLambda, testARN, "")

	if err != nil {
		panic(err)
	}

	return mockLambda, awsLambda
}

func getMockLambdaAlias() (*Lambda, *mocks.LambdaAPI) {
	testARN := "arn:aws:test"
	testArchitecture := "x86_64"
	testMemorySize := 256

	awsLambda := new(mocks.LambdaAPI)
	awsLambda.On("GetFunctionConfiguration", mock.AnythingOfType("*lambda.GetFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(testArchitecture)},
		MemorySize:    aws.Int64(int64(testMemorySize)),
	}, nil)

	awsLambda.On("GetAlias", mock.AnythingOfType("*lambda.GetAliasInput")).Return(&lambda.AliasConfiguration{
		FunctionVersion: aws.String("1"),
	}, nil)

	mockLambda, err := NewLambda(awsLambda, testARN, "live")

	if err != nil {
		panic(err)
	}

	return mockLambda, awsLambda
}

func TestChangeMemory(t *testing.T) {
	mockLambda, awsLambda := getMockLambda()

	awsLambda.On("UpdateFunctionConfiguration", mock.AnythingOfType("*lambda.UpdateFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(mockLambda.Architecture)},
		MemorySize:    aws.Int64(int64(512)),
	}, nil)

	err := mockLambda.ChangeMemory(512)
	assert.Nil(t, err, "Expected no error")

	awsLambda.AssertCalled(t, "UpdateFunctionConfiguration", &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String(mockLambda.Arn),
		MemorySize:   aws.Int64(512),
	})
}

func TestChangeMemoryError(t *testing.T) {
	mockLambda, awsLambda := getMockLambda()

	awsLambda.On("UpdateFunctionConfiguration", mock.AnythingOfType("*lambda.UpdateFunctionConfigurationInput")).Return(nil, errors.New("test"))

	err := mockLambda.ChangeMemory(512)
	assert.NotNil(t, err, "Expected error")
	assert.Equal(t, "test", err.Error())
}

func TestChangeMemoryResourceConflictError(t *testing.T) {
	mockLambda, awsLambda := getMockLambda()

	awsLambda.On("UpdateFunctionConfiguration", mock.AnythingOfType("*lambda.UpdateFunctionConfigurationInput")).Return(nil, awserr.New(lambda.ErrCodeResourceConflictException, "ResourceConflictException message", nil))

	err := mockLambda.ChangeMemory(512)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, errors.New("maximum number of retries to change Lambda memory exceeded"), err)
}

func TestChangeMemoryAlias(t *testing.T) {
	mockLambda, awsLambda := getMockLambdaAlias()
	newVersion := "1337"

	awsLambda.On("UpdateFunctionConfiguration", mock.AnythingOfType("*lambda.UpdateFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(mockLambda.Architecture)},
		MemorySize:    aws.Int64(int64(512)),
	}, nil)

	awsLambda.On("PublishVersion", mock.AnythingOfType("*lambda.PublishVersionInput")).Return(&lambda.FunctionConfiguration{
		Version: aws.String(newVersion),
	}, nil)

	awsLambda.On("UpdateAlias", mock.AnythingOfType("*lambda.UpdateAliasInput")).Return(&lambda.AliasConfiguration{
		FunctionVersion: aws.String(newVersion),
	}, nil)

	err := mockLambda.ChangeMemory(512)
	assert.Nil(t, err, "Expected no error")

	awsLambda.AssertCalled(t, "UpdateFunctionConfiguration", &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String(mockLambda.Arn),
		MemorySize:   aws.Int64(512),
	})

	awsLambda.AssertCalled(t, "UpdateAlias", &lambda.UpdateAliasInput{
		FunctionName:    aws.String(mockLambda.Arn),
		FunctionVersion: aws.String(newVersion),
		Name:            aws.String(mockLambda.Alias),
	})
}

func TestChangeMemoryAliasErrorCreatingNewVersion(t *testing.T) {
	mockLambda, awsLambda := getMockLambdaAlias()

	awsLambda.On("UpdateFunctionConfiguration", mock.AnythingOfType("*lambda.UpdateFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(mockLambda.Architecture)},
		MemorySize:    aws.Int64(int64(512)),
	}, nil)

	awsLambda.On("PublishVersion", mock.AnythingOfType("*lambda.PublishVersionInput")).Return(nil, errors.New("test"))

	err := mockLambda.ChangeMemory(512)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, "test", err.Error())
}

func TestChangeMemoryAliasErrorUpdatingAlias(t *testing.T) {
	mockLambda, awsLambda := getMockLambdaAlias()
	newVersion := "1337"

	awsLambda.On("UpdateFunctionConfiguration", mock.AnythingOfType("*lambda.UpdateFunctionConfigurationInput")).Return(&lambda.FunctionConfiguration{
		Architectures: []*string{aws.String(mockLambda.Architecture)},
		MemorySize:    aws.Int64(int64(512)),
	}, nil)

	awsLambda.On("PublishVersion", mock.AnythingOfType("*lambda.PublishVersionInput")).Return(&lambda.FunctionConfiguration{
		Version: aws.String(newVersion),
	}, nil)

	awsLambda.On("UpdateAlias", mock.AnythingOfType("*lambda.UpdateAliasInput")).Return(nil, errors.New("test"))

	err := mockLambda.ChangeMemory(512)
	assert.Error(t, err, "Expected error")
	assert.Equal(t, "test", err.Error())
}
