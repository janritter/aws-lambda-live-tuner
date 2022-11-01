# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class AwsLambdaLiveTuner < Formula
  desc "Tool to optimize Lambda functions on real incoming events"
  homepage "https://github.com/janritter/aws-lambda-live-tuner"
  version "0.7.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/janritter/aws-lambda-live-tuner/releases/download/0.7.1/aws-lambda-live-tuner_0.7.1_darwin_arm64.tar.gz"
      sha256 "56203f3eb93cc75295d0ab4ee27033a6ce1fd05c0024b9bcdc30545636ee1e75"

      def install
        bin.install "aws-lambda-live-tuner"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/janritter/aws-lambda-live-tuner/releases/download/0.7.1/aws-lambda-live-tuner_0.7.1_darwin_amd64.tar.gz"
      sha256 "070bb29b4489ed7367b6cd385eb1cca2f3c30a913863642e1c1e3e57dc7a9ff1"

      def install
        bin.install "aws-lambda-live-tuner"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/janritter/aws-lambda-live-tuner/releases/download/0.7.1/aws-lambda-live-tuner_0.7.1_linux_arm64.tar.gz"
      sha256 "e1a65a28300ea1f80aafabe53f1a94a7101d540d392e83a286100f654984f4a0"

      def install
        bin.install "aws-lambda-live-tuner"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/janritter/aws-lambda-live-tuner/releases/download/0.7.1/aws-lambda-live-tuner_0.7.1_linux_amd64.tar.gz"
      sha256 "d27ab6ac28686bfe0be5840be6c8fdb0a1ac85f4cbbcc6fc86bebeba7d90d6dc"

      def install
        bin.install "aws-lambda-live-tuner"
      end
    end
  end
end
