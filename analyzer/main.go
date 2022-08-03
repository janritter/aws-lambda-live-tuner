package analyzer

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"go.uber.org/zap"
)

type AnalyzerAPI interface {
	CheckInvocations(lambdaARN string, memory int) error
}

type Analyzer struct {
	cloudwatch cloudwatchlogsiface.CloudWatchLogsAPI
	logger     *zap.SugaredLogger
}

func NewAnalyzer(cloudwatch cloudwatchlogsiface.CloudWatchLogsAPI, logger *zap.SugaredLogger) *Analyzer {
	return &Analyzer{
		cloudwatch: cloudwatch,
		logger:     logger,
	}
}
