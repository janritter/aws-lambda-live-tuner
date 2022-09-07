# [0.4.0](https://github.com/janritter/aws-lambda-live-tuner/compare/0.3.0...0.4.0) (2022-09-07)


### Bug Fixes

* **deps:** go mod updates ([e9efc59](https://github.com/janritter/aws-lambda-live-tuner/commit/e9efc59031f343a22536ab460049ecb5d798aa88))


### Features

* **#16:** added support for arm architecture in cost calculation ([803be32](https://github.com/janritter/aws-lambda-live-tuner/commit/803be321bfe9e3496132cd508163ce44e912552c)), closes [#16](https://github.com/janritter/aws-lambda-live-tuner/issues/16)

# [0.3.0](https://github.com/janritter/aws-lambda-live-tuner/compare/0.2.4...0.3.0) (2022-09-01)


### Features

* added csv output option ([046e259](https://github.com/janritter/aws-lambda-live-tuner/commit/046e259a8a52e81982c9cb27f04f1d21d268f1e8))

## [0.2.4](https://github.com/janritter/aws-lambda-live-tuner/compare/0.2.3...0.2.4) (2022-08-29)


### Bug Fixes

* **deps:** update minor - go ([9098721](https://github.com/janritter/aws-lambda-live-tuner/commit/909872164e53b2641ec8acb3e2facefa4bbbdbe2))

## [0.2.3](https://github.com/janritter/aws-lambda-live-tuner/compare/0.2.2...0.2.3) (2022-08-26)


### Bug Fixes

* **deps:** update minor - go ([4c283c2](https://github.com/janritter/aws-lambda-live-tuner/commit/4c283c2085b0101b2aa20f61e4067238f4e19583))

## [0.2.2](https://github.com/janritter/aws-lambda-live-tuner/compare/0.2.1...0.2.2) (2022-08-14)


### Bug Fixes

* aws-sdk package and go version updates ([cc552a8](https://github.com/janritter/aws-lambda-live-tuner/commit/cc552a8b4d5fc94745dafc449022f68f29aaf082))

## [0.2.1](https://github.com/janritter/aws-lambda-live-tuner/compare/0.2.0...0.2.1) (2022-08-04)


### Bug Fixes

* removed wrong increment help description ([07721cc](https://github.com/janritter/aws-lambda-live-tuner/commit/07721cceaa2f5c940055a8e9f3080f5af9158bab))

# [0.2.0](https://github.com/janritter/aws-lambda-live-tuner/compare/0.1.1...0.2.0) (2022-08-04)


### Bug Fixes

* include error in formatting string ([53675cc](https://github.com/janritter/aws-lambda-live-tuner/commit/53675cc6e2c3fa4f0fa5a84edd307c1037169cf4))


### Features

* added additional validations and adapted the existing ones to reflect the correct limitations by AWS - resolves [#4](https://github.com/janritter/aws-lambda-live-tuner/issues/4) ([b6c85e3](https://github.com/janritter/aws-lambda-live-tuner/commit/b6c85e30a5b18d4de85f59a88e1049704971d267))
* made cloudwatch time window dynamic based on wait time ([84925f5](https://github.com/janritter/aws-lambda-live-tuner/commit/84925f54c8733e6e111f3a6ba52d26e110b7b51b))
* replaced zap logger with colored logging to improve cli experience ([75a3c12](https://github.com/janritter/aws-lambda-live-tuner/commit/75a3c1220dacf0625b320510584ab779f7ab3277))
* sort final output list by memory ([e5fff7d](https://github.com/janritter/aws-lambda-live-tuner/commit/e5fff7dba153216e7b387c6ca7142dec7c5502f3))

## [0.1.1](https://github.com/janritter/aws-lambda-live-tuner/compare/0.1.0...0.1.1) (2022-08-04)


### Bug Fixes

* reset lambda memory config to the pre test value - resolves [#2](https://github.com/janritter/aws-lambda-live-tuner/issues/2) ([9fac8ab](https://github.com/janritter/aws-lambda-live-tuner/commit/9fac8aba7010bea345f0ff21f44fa7a0c2261488))

# [0.1.0](https://github.com/janritter/aws-lambda-live-tuner/compare/0.0.0...0.1.0) (2022-08-04)


### Features

* added calculation of cost ([e8a2259](https://github.com/janritter/aws-lambda-live-tuner/commit/e8a225995596c227491c12fcd0548a1ceffb950c))
* added cloudwatch log insights analyzer and complete main test loop ([2d275c5](https://github.com/janritter/aws-lambda-live-tuner/commit/2d275c52829e8d0f04f0329888ec7d589901583a))
* added memory changer and first code for the main test loop ([cc6185c](https://github.com/janritter/aws-lambda-live-tuner/commit/cc6185c8d4be390c898ba9cc806a5b5273943501))
