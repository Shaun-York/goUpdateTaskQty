# goUpdateTaskQty - NetSuite WorkOrder Completion AWS Lambda 1 of 3

Consume SQS messages from input queue. Update completed qty via GraphQL server, delete message, pass to next queue.

AWSAPI Gateway -> SQS -> [goUpdateTaskQty](https://github.com/Shaun-York/goUpdateTaskQty) -> SQS -> [goUpdateNetSuite](https://github.com/Shaun-York/goUpdateNetSuite) -> SQS -> [goUpdateCompletion](https://github.com/Shaun-York/goUpdateCompletion)
