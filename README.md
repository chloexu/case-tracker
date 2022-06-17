# USCIS Case Tracker

## About
I created this mini program to check my H1B case transfer status hourly and send message to a Slack webhook.

When the case status changes, as in being different from the `DEFAULT_STATUS` specified in config json file, the slackbot will `@here` the Slack channel.


## Usage

I configured the code to run on AWS Lambda, with a hourly cron job configured on AWS Cloudwatch triggering the Lambda function. `deploy.sh` builds the file needed to upload to AWS Lambda.

To run the code in local, see command in `dev_start.sh`

## Environment variables

Make two copies of `config/placeholder.json` in `config` directory and name them as below:
* `config/default.json` to store dev / local env vars. [You need this file to run the code locally.]
* `config/prod.json` to store prod env vars. [You need to this file to build and deploy to Lambda.]
