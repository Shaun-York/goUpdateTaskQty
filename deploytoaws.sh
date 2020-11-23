# aws-cli needs to be install and authenticated 
# AWS Account number
AWS_ACCOUNT=""
# AWS role with proper privileges 
AWS_ROLE=""
GOARCH="amd64"
GOOS="linux"

GOARCH=amd64 GOOS=linux go build . -ldflags="-s -w"

zip -r goUpdateTaskQty.zip goUpdateTaskQty

aws lambda create-function \
    --function-name goUpdateTaskQty \
    --runtime go1.x \
    --zip-file fileb://goUpdateTaskQty.zip \
    --handler goUpdateTaskQty \
    --role arn:aws:iam::$AWS_ACCOUNT:role/service-role/$AWS_ROLE