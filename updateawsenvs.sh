# AWS SQS URL that feeds this function (AWS APIGATEWAY)
C_INPUT_QUEUE=""
# AWS SQS URL to send successful payloads (NetSuite Update Queue) to.
C_OUTPUT_QUEUE=""
# URL to Graphql server
C_GQL_SERVER_URL=""
# Update table w/ gql server retries (2 should be fine for same VPC)
C_GQL_SERVER_RETRYS=""
# GQL server secret 
#TODO setup jwts
C_GQL_SERVER_SECRET=""

ENVS="{\
INPUT_QUEUE=$C_INPUT_QUEUE,\
OUTPUT_QUEUE=$C_OUTPUT_QUEUE,\
GQL_SERVER_URL=$C_GQL_SERVER_URL,\
GQL_SERVER_RETRYS=$C_GQL_SERVER_RETRYS,\
GQL_SERVER_SECRET=$C_GQL_SERVER_SECRET\
}"

aws lambda update-function-configuration \
    --function-name  goUpdateTaskQty \
    --environment Variables=$ENVS