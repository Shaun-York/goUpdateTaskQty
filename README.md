# goUpdateTaskQty - NetSuite WorkOrder Completion AWS Lambda 1 of 3

Consume SQS messages from input queue. Update completed qty via GraphQL server, delete message, pass to next queue.

AWSAPI Gateway -> SQS -> [goUpdateTaskQty](https://github.com/Shaun-York/goUpdateTaskQty) -> SQS -> [goUpdateNetSuite](https://github.com/Shaun-York/goUpdateNetSuite) -> SQS -> [goUpdateCompletion](https://github.com/Shaun-York/goUpdateCompletion)

## Payload - Completion table columns

```json
{
    "workordercompletion_id": "to be filled workorder completion internailid",
    "operation_sequence": "NetSuite operation sequence",
    "last_completion": "boolean this completion completes build",
    "remaining_qty": "input qty - completed",
    "mfgoptask_id": "NetSuite mfg operation task internalid",
    "completedQty": "Qty completed",
    "workorder_id": "NetSuite workorder internalid",
    "operator_id": "Operators",
    "worktime_id": "Worktime table pk",
    "location_id": "Location completed at internalid",
    "machine_id": "Machine completed on pk",
    "updated_at": "updated at",
    "created_at": "created at",
    "workcenter": "NetSuite entity group",
    "scrapQty": "Qty scrapped in completion",
    "item_id": "NetSuite item internalid",
    "action": "mk or rm",
    "id": "table pk",
}
```
