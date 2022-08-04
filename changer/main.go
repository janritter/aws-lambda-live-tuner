package changer

import (
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
)

type ChangerAPI interface {
	ChangeMemory(lambdaARN string, memory int) error
	GetCurrentMemoryValue(lambdaARN string) (int, error)
}

type Changer struct {
	lambda lambdaiface.LambdaAPI
}

func NewChanger(lambda lambdaiface.LambdaAPI) *Changer {
	return &Changer{
		lambda: lambda,
	}
}
