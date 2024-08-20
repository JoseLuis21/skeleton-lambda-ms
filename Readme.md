# Skeleton ms

## Environment Variables

To run this project, you will need to add the following environment variables to your environment

```bash
TEST_ENV=
```

## Run Locally

## Install SAM

- [SAM AWS](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)

## Build and start lambda

```bash
  sam build && sam local invoke TestGolangLocalSqs --event sqs_event.json
```
