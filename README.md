# diomedes-search

Get a notification on Telegram whenever your movie opens bookings in a theater of your choice.

## Pre-requisites

1. [Install AWS CLI (v2) by following these instructions](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html).
2. [Install SAM CLI by following these instructions](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html).
3. [Go 1.17+](https://go.dev/dl/)

## Build

```
mkdir build
go build -o ./build ./cmd/...
```

## Deploy

1. Run `aws configure` and add AWS Access, Secret Keys, and select your default region and response format (json).

   [Learn how to create these keys here.](https://docs.aws.amazon.com/IAM/latest/UserGuide/getting-started_create-admin-group.html)

2. `sam deploy --guided`

   Follow the instructions, and check [AWS Lambda Console](https://console.aws.amazon.com/lambda)

## Post-deploy

1. Go to the Notification function on AWS Console and set the `TG_TOKEN` to your Telegram bot token.
2. Add an alert in `MovieTable` like this (I use the AWS Console to keep it simple):
![image](https://user-images.githubusercontent.com/13501594/144083714-c4ca920f-de45-4819-ad5d-6ea434511661.png)


# Contributing

Open an Issue for any help, reporting bugs.
Open a PR for enhancements and bugfixes.

# License

[MIT License](LICENSE)
