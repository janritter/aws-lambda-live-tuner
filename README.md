# AWS Lambda Live Tuner

[![CircleCI](https://circleci.com/gh/janritter/aws-lambda-live-tuner/tree/main.svg?style=svg)](https://circleci.com/gh/janritter/aws-lambda-live-tuner/tree/main)

> **Warning**
> AWS Lambda Live Tuner is still in very early development, functionality might change between releases until version 1.0.0.
> Ideas and feedback are very welcome

AWS Lambda Live Tuner tests memory configurations based on real incoming events instead of a single test event.

Let's imagine we are testing a Lambda function that processes a queue, since the Lambda function is idempotent, messages that have already been processed will be successfully processed again. Using the same test event on all invocations might falsify the results because all subsequent invocations after the initial one might be way faster (event was already processed before). Using different incoming events instead helps you test the actual behavior of the Lambda.

This project is heavily inspried by the open source tool [aws-lambda-power-tuning](https://github.com/alexcasalboni/aws-lambda-power-tuning) 

## Prerequisites

- Configured AWS credentials

## Installation via Homebrew (For Mac / Linux)

### Get the formula

```bash
brew tap janritter/aws-lambda-live-tuner https://github.com/janritter/aws-lambda-live-tuner
```

### Install formula

```bash
brew install aws-lambda-live-tuner
```

## Usage

```text
aws-lambda-live-tuner --help

Usage:
  aws-lambda-live-tuner [flags]

Flags:
      --config string            config file (default is $HOME/.aws-lambda-live-tuner.yaml)
  -h, --help                     help for aws-lambda-live-tuner
      --lambda-arn string        ARN of the Lambda function to optimize
      --memory-increment int     Increments for the memory configuration added to the min value until the max value is reached (default 64)
      --memory-max int           Upper memory limit for the optimization (default 2048)
      --memory-min int           Lower memory limit for the optimization (default 128)
      --min-requests int         Minimum number of requests the Lambda function must receive before continuing with the next memory configuration (default 5)
      --output-filename string   Filename for the output csv
      --wait-time int            Wait time in seconds between CloudWatch Log insights queries (default 180)
```

### Examples

```bash
aws-lambda-live-tuner --lambda-arn arn:aws:lambda:eu-central-1:1234567890:function:my-lambda-name
```

## License and Author

Author: Jan Ritter

License: MIT
