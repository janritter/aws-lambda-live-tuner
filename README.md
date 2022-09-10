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

### Check help

```text
aws-lambda-live-tuner --help
```

### Test Lambda

```bash
aws-lambda-live-tuner --lambda-arn arn:aws:lambda:eu-central-1:1234567890:function:my-lambda-name
```

## Limitations

- Lambda@Edge functions are currently not supported
- Lambda tiered pricing is not considered
  - Because we can't know in which Lambda pricing tier you are operating, we will always use Tier 1. Since higher tiers just reduce the GB/second price, the lowest price results will still be valid for you.

## License and Author

Author: Jan Ritter

License: MIT
