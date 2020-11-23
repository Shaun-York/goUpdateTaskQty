# aws-cli needs to be install and authenticated 
# AWS Account number
AWS_ACCOUNT=""
# AWS role with proper privileges 
AWS_ROLE=""
GOARCH="amd64"
GOOS="linux"

GOARCH=amd64 GOOS=linux go build -ldflags="-s -w"

zip -r goUpdateTaskQty.zip goUpdateTaskQty

aws lambda update-function-code \
    --function-name goUpdateTaskQty \
    --zip-file fileb://goUpdateTaskQty.zip