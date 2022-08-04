package changer

import (
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"go.uber.org/zap"
)

type ChangerAPI interface {
	ChangeMemory(memory int) error
}

type Changer struct {
	lambda lambdaiface.LambdaAPI
	logger *zap.SugaredLogger
}

func NewChanger(lambda lambdaiface.LambdaAPI, logger *zap.SugaredLogger) *Changer {
	return &Changer{
		lambda: lambda,
		logger: logger,
	}
}
