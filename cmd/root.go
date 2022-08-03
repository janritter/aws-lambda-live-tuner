package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/janritter/aws-lambda-live-tuner/changer"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"go.uber.org/zap"
)

var cfgFile string
var minRequests int
var memoryMin int
var memoryMax int
var memoryIncrement int
var lambdaARN string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-lambda-live-tuner",
	Short: "Tool to optimize Lambda functions on real incoming events",
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewDevelopment()
		defer logger.Sync()
		sugaredLogger := logger.Sugar()

		sugaredLogger.Info("Starting AWS Lambda Live Tuner")
		validateInputs()

		awsSession := session.Must(session.NewSession())
		lambdaSvc := lambda.New(awsSession)

		changer := changer.NewChanger(lambdaSvc, sugaredLogger)

		for memory := memoryMin; memory <= memoryMax; memory += memoryIncrement {
			sugaredLogger.Infof("Starting test for %dMB", memory)

			err := changer.ChangeMemory(lambdaARN, memory)
			if err != nil {
				os.Exit(1)
			}
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
	rootCmd.PersistentFlags().IntVar(&memoryIncrement, "memory-increment", 64, "Increments for the memory configuration added to the min value until the max value is reached. The increment must be a multiple of 64")
	rootCmd.PersistentFlags().StringVar(&lambdaARN, "lambda-arn", "", "ARN of the Lambda function to optimize")
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
	validateMemoryMinValue()
	validateMemoryMaxValue()
	validateMemoryIncrement()
	validateMinRequests()
}

func validateMemoryMinValue() {
	if memoryMin < 128 {
		log.Println("Memory min value must be greater than or equal to 128")
		os.Exit(1)
	}
	if memoryMin%64 != 0 {
		log.Println("Memory min value must be a multiple of 64 with the minimal value of 128")
		os.Exit(1)
	}
}

func validateMemoryMaxValue() {
	if memoryMax <= memoryMin {
		log.Println("Memory max value must be greater than the min value")
		os.Exit(1)
	}
	if memoryMax%64 != 0 {
		log.Println("Memory max value must be a multiple of 64 with the minimal value of 192")
		os.Exit(1)
	}
}

func validateMemoryIncrement() {
	if memoryIncrement%64 != 0 {
		log.Println("Memory increment value must be a multiple of 64")
		os.Exit(1)
	}
}

func validateMinRequests() {
	if minRequests < 1 {
		log.Println("Minimum number of requests must be greater than 0")
		os.Exit(1)
	}
}
