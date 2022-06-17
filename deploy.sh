export PATH_TO_CONFIG=./config/prod.json

GOARCH=amd64 GOOS=linux go build .

zip package.zip case-tracker config/*

echo "Now you can upload package.zip to your Lambda function!"

rm -r -f case-tracker
