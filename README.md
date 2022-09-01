# AWS Lambda Live Tuner

[![CircleCI](https://circleci.com/gh/janritter/aws-lambda-live-tuner/tree/main.svg?style=svg)](https://circleci.com/gh/janritter/aws-lambda-live-tuner/tree/main)

> Tool to optimize Lambda functions on real incoming events

## Prerequisites

- TODO

## Installation via go

### Clone git repo

```bash
git clone git@github.com:janritter/aws-lambda-live-tuner.git
```

### Open project directory

```bash
cd aws-lambda-live-tuner
```

### Install via go

```bash
go install
```

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
