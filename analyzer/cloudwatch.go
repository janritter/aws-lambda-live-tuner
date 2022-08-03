package analyzer

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"go.uber.org/zap"
)

func (a *Analyzer) CheckInvocations(lambdaARN string, memory int) (map[string]float64, error) {
	functionName := getFunctionNameFromARN(lambdaARN)
	logGroupName := fmt.Sprintf("/aws/lambda/%s", functionName)

	output, err := a.cloudwatch.StartQuery(&cloudwatchlogs.StartQueryInput{
		QueryString:  aws.String(fmt.Sprintf(`filter @type = "REPORT" and @message like "Memory Size: %d MB"`, memory)),
		LogGroupName: aws.String(logGroupName),
		StartTime:    aws.Int64(time.Now().Add(-5 * time.Minute).Unix()),
		EndTime:      aws.Int64(time.Now().Unix()),
	})
	if err != nil {
		a.logger.Error("Failed starting CloudWatch log insights query: ", zap.Error(err))
		return nil, err
	}

	queryID := *output.QueryId

	resultMap := make(map[string]float64)

	for {
		queryResultOutput, err := a.cloudwatch.GetQueryResults(&cloudwatchlogs.GetQueryResultsInput{
			QueryId: aws.String(queryID),
		})
		if err != nil {
			a.logger.Error("Failed getting CloudWatch log insights query results: ", zap.Error(err))
			return nil, err
		}

		if *queryResultOutput.Status == cloudwatchlogs.QueryStatusComplete {
			results := queryResultOutput.Results

			for _, fields := range results {
				for _, field := range fields {
					if *field.Field == "@message" {
						// log.Println(*field.Value)
						id, duration, err := getDurationWithRequestIdFromMessage(*field.Value)
						if err != nil {
							a.logger.Error(zap.Error(err))
							return nil, err
						}
						resultMap[id] = duration
					}
				}
			}

			a.logger.Info("CloudWatch log insights query completed successfully")
			break
		}

		if !(*queryResultOutput.Status == cloudwatchlogs.QueryStatusComplete || *queryResultOutput.Status == cloudwatchlogs.QueryStatusRunning || *queryResultOutput.Status == cloudwatchlogs.QueryStatusScheduled) {
			a.logger.Error("CloudWatch log insights query is not an expected status: ", zap.String("status", *queryResultOutput.Status))
			return nil, fmt.Errorf("CloudWatch log insights query is not an expected status: %s", *queryResultOutput.Status)
		}
	}

	return resultMap, nil
}

func getFunctionNameFromARN(arn string) string {
	elements := strings.Split(arn, ":")
	return elements[len(elements)-1]
}

func getDurationWithRequestIdFromMessage(message string) (string, float64, error) {
	requestId, err := getRequestIdFromMessage(message)
	if err != nil {
		return "", -1, err
	}

	duration, err := getDurationFromMessage(message)
	if err != nil {
		return "", -1, err
	}

	return requestId, duration, nil
}

func getRequestIdFromMessage(message string) (string, error) {
	start := strings.Index(message, "RequestId: ")
	end := strings.Index(message, "\tDuration")

	if start == -1 || end == -1 {
		return "", fmt.Errorf("Failed to parse RequestId from message: %s", message)
	}

	return (message[start+11 : end]), nil
}

func getDurationFromMessage(message string) (float64, error) {
	start := strings.Index(message, "Duration: ")
	end := strings.Index(message, " ms")

	if start == -1 || end == -1 {
		return -1, fmt.Errorf("Failed to parse duration from message: %s", message)
	}

	duration := (message[start+10 : end])

	durationFloat, err := strconv.ParseFloat(duration, 64)
	if err != nil {
		return -1, fmt.Errorf("Failed to convert duration to int: %s", duration)
	}
	return durationFloat, nil
}
