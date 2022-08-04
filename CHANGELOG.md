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
