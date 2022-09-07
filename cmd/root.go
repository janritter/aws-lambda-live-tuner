package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/janritter/aws-lambda-live-tuner/analyzer"
	"github.com/janritter/aws-lambda-live-tuner/changer"
	"github.com/janritter/aws-lambda-live-tuner/cost"
	"github.com/janritter/aws-lambda-live-tuner/helper"
	"github.com/janritter/aws-lambda-live-tuner/output"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"golang.org/x/exp/maps"
)

var cfgFile string
var minRequests int
var memoryMin int
var memoryMax int
var waitTime int
var memoryIncrement int
var lambdaARN string
var outputFilename string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-lambda-live-tuner",
	Short: "Tool to optimize Lambda functions on real incoming events",
	Run: func(cmd *cobra.Command, args []string) {
		helper.LogSuccess("Starting AWS Lambda Live Tuner")
		validateInputs()

		awsSession := session.Must(session.NewSession())
		lambdaSvc := lambda.New(awsSession)
		cloudwatchlogsSvc := cloudwatchlogs.New(awsSession)

		changer := changer.NewChanger(lambdaSvc)
		analyzer := analyzer.NewAnalyzer(cloudwatchlogsSvc, waitTime)

		resetMemoryValue, err := changer.GetCurrentMemoryValue(lambdaARN)
		if err != nil {
			os.Exit(1)
		}
		helper.LogInfo("Memory value before test start: %d", resetMemoryValue)

		architecture, err := changer.GetArchitecture(lambdaARN)
		if err != nil {
			os.Exit(1)
		}
		helper.LogInfo("Architecture of Lambda: %s", architecture)

		durationResults := make(map[int]float64)
		costResults := make(map[int]float64)
		for memory := memoryMin; memory <= memoryMax; memory += memoryIncrement {
			helper.LogInfo("Starting test for %dMB", memory)

			err := changer.ChangeMemory(lambdaARN, memory)
			if err != nil {
				changer.ChangeMemory(lambdaARN, resetMemoryValue)
				os.Exit(1)
			}

			invocations := make(map[string]float64)
			for len(invocations) < minRequests {
				newInvocations, err := analyzer.CheckInvocations(lambdaARN, memory)
				if err != nil {
					changer.ChangeMemory(lambdaARN, resetMemoryValue)
					os.Exit(1)
				}

				maps.Copy(invocations, newInvocations)

				helper.LogInfo("Total number of invocations analyzed for memory config: %d", len(invocations))
				if len(invocations) > minRequests {
					break
				}

				helper.LogNotice("Waiting %d seconds before next analysis", waitTime)
				time.Sleep(time.Duration(waitTime) * time.Second)
			}

			helper.LogNotice("Calculating average duration for %dMB memory", memory)
			average := calculateAverageOfMap(invocations)
			durationResults[memory] = average

			costResult := cost.Calculate(average, memory, architecture, getRegionFromARN(lambdaARN))

			costResults[memory] = costResult

			helper.LogSuccess("[RESULT] Memory: %d MB - Average Duration: %f ms - Cost %.10f USD", memory, average, costResult)

			helper.LogInfo("Test for %dMB finished", memory)
		}

		csvRecords := [][]string{
			{"memory", "duration", "cost"},
		}
		sorted := memorySortedList(durationResults)
		for _, memory := range sorted {
			helper.LogSuccess("%d MB - Average Duration: %f ms - Cost: %.10f USD", memory, durationResults[memory], costResults[memory])
			csvRecords = append(csvRecords, []string{
				fmt.Sprint(memory), fmt.Sprintf("%f", durationResults[memory]), fmt.Sprintf("%.10f", costResults[memory]),
			})
		}

		if outputFilename != "" {
			output.WriteCSV(outputFilename, csvRecords)
		}

		helper.LogInfo("Changing Lambda to pre-test memory value of %dMB", resetMemoryValue)
		err = changer.ChangeMemory(lambdaARN, resetMemoryValue)
		if err != nil {
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aws-lambda-live-tuner.yaml)")

	rootCmd.PersistentFlags().IntVar(&minRequests, "min-requests", 5, "Minimum number of requests the Lambda function must receive before continuing with the next memory configuration")
	rootCmd.PersistentFlags().IntVar(&memoryMin, "memory-min", 128, "Lower memory limit for the optimization")
	rootCmd.PersistentFlags().IntVar(&memoryMax, "memory-max", 2048, "Upper memory limit for the optimization")
	rootCmd.PersistentFlags().IntVar(&memoryIncrement, "memory-increment", 64, "Increments for the memory configuration added to the min value until the max value is reached")
	rootCmd.PersistentFlags().StringVar(&lambdaARN, "lambda-arn", "", "ARN of the Lambda function to optimize")
	rootCmd.PersistentFlags().IntVar(&waitTime, "wait-time", 180, "Wait time in seconds between CloudWatch Log insights queries")
	rootCmd.PersistentFlags().StringVar(&outputFilename, "output-filename", "", "Filename for the output csv")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".aws-lambda-live-tuner" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".aws-lambda-live-tuner")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func validateInputs() {
	validateLambdaARN()
	validateWaitTime()
	validateMemoryMinValue()
	validateMemoryMaxValue()
	validateMemoryIncrement()
	validateMinRequests()
}

func validateWaitTime() {
	if waitTime < 30 {
		helper.LogError("Wait time must be at least 30 seconds")
		os.Exit(1)
	}
}

func validateLambdaARN() {
	if !strings.HasPrefix(lambdaARN, "arn:aws:lambda:") {
		helper.LogError("Lambda ARN must be in the format arn:aws:lambda:<region>:<account-id>:function:<function-name>")
		os.Exit(1)
	}
}

func getRegionFromARN(arn string) string {
	elements := strings.Split(arn, ":")
	return elements[3]
}

func validateMemoryMinValue() {
	if memoryMin < 128 {
		helper.LogError("Memory min value must be greater than or equal to 128")
		os.Exit(1)
	}
}

func validateMemoryMaxValue() {
	if memoryMax <= memoryMin {
		helper.LogError("Memory max value must be greater than the min value")
		os.Exit(1)
	}
	if memoryMax > 10240 {
		helper.LogError("Memory max value must be less than or equal to 10240")
		os.Exit(1)
	}
}

func validateMemoryIncrement() {
	if memoryIncrement < 1 {
		helper.LogError("Memory increment must be greater than 0")
		os.Exit(1)
	}
}

func validateMinRequests() {
	if minRequests < 1 {
		log.Println("Minimum number of requests must be greater than 0")
		os.Exit(1)
	}
}

func calculateAverageOfMap(data map[string]float64) float64 {
	var total float64 = 0.0
	for _, value := range data {
		total += value
	}
	return total / float64(len(data))
}

func memorySortedList(results map[int]float64) []int {
	keys := make([]int, 0, len(results))
	for key := range results {
		keys = append(keys, key)
	}

	sort.IntSlice(keys).Sort()

	return keys
}
