# AWS Lambda Live Tuner

[![CircleCI](https://circleci.com/gh/janritter/aws-lambda-live-tuner/tree/main.svg?style=svg)](https://circleci.com/gh/janritter/aws-lambda-live-tuner/tree/main)

> **Warning**
> AWS Lambda Live Tuner is still in very early development, functionality might change between releases until version 1.0.0.
> Ideas and feedback are very welcome

AWS Lambda Live Tuner tests memory configurations based on real incoming events instead of a single test event.

Let's imagine we are testing a Lambda function that processes a queue, since the Lambda function is idempotent, messages that have already been processed will be successfully processed again. Using the same test event on all invocations might falsify the results because all subsequent invocations after the initial one might be way faster (event was already processed before). Using different incoming events instead helps you test the actual behavior of the Lambda.

This project is heavily inspired by the open source tool [aws-lambda-power-tuning](https://github.com/alexcasalboni/aws-lambda-power-tuning) 

## Prerequisites

- Configured AWS credentials

## Installation

### Via Homebrew (For Mac / Linux)

#### Get the formula

```bash
brew tap janritter/aws-lambda-live-tuner https://github.com/janritter/aws-lambda-live-tuner
```

#### Install formula

```bash
brew install aws-lambda-live-tuner
```

### Via download of pre-build binaries

1. Open the [latest release page](https://github.com/janritter/aws-lambda-live-tuner/releases/latest)
2. Download the archive with the pre-build binary for your operating system and architecture
    - For Linux with amd64 architecture this would be `aws-lambda-live-tuner_<version>_linux_amd64.tar.gz`
3. Unzip the downloaded archive
4. Start using aws-lambda-live-tuner

### Via local build

This option requires go to be installed

#### Clone the repo

```bash
git clone git@github.com:janritter/aws-lambda-live-tuner.git
```

#### Build

```make
make build
```

The binary is saved in `bin` inside the project folder

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
- Lambda function aliases are currently not supported
- Lambda tiered pricing is not considered
  - Because we can't know in which Lambda pricing tier you are operating, we will always use Tier 1. Since higher tiers just reduce the GB/second price, the lowest price results will still be valid for you.

## Development

### Regenerate AWS SDK mocks for testing

```bash
mockery
```

## License and Author

Author: Jan Ritter

License: MIT
