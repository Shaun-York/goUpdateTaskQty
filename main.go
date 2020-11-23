package main

import (
	"context"
	"encoding/json"
	C "goUpdateTaskQty/completion"
	Gql "goUpdateTaskQty/gql"
	S "goUpdateTaskQty/sqssrv"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type tableStruct struct {
	Name string `json:"name"`
}

type eventData struct {
	Old C.Completion `json:"old"`
	New C.Completion `json:"new"`
}

type eventStruct struct {
	Operation string    `json:"op"`
	Data      eventData `json:"data"`
}

type hasuraEvent struct {
	Table *tableStruct `json:"table"`
	Event *eventStruct `json:"event"`
	Op    string       `json:"op"`
}
type payloadStruct struct {
    Payload *hasuraEvent `json:"payload"`
}
// Handle update task
func Handle(ctx context.Context, event events.SQSEvent) (string, error) {
    var msgID string
    var failed error

    for _, v := range event.Records {
        recid := v.ReceiptHandle
        msgID = v.MessageId

        hevent := &payloadStruct{}

        if herr := json.Unmarshal([]byte(v.Body), hevent); herr != nil {
            log.Printf("Unmarshal into hevet failed pl: %s", v.Body)
            failed = herr
            break
        }

        compl := &hevent.Payload.Event.Data.New
        compl.ReceiptHandle = recid


        uperr := Gql.UpdateTaskQty(compl)

        if uperr != nil {
            failed = uperr
            break
        }

        srv := S.CompletionsToNetSuite{}
        sqserr := srv.GetSrv()

        if sqserr != nil {
            failed = sqserr
            break
        }

        sqsmsg, sederr := compl.SqsMsg()

        if sederr != nil {
            log.Printf("compl.SqsMsg() failed pl")
            failed = sederr
            break
        }

        senderr := srv.Send(sqsmsg, compl.SqsDelMsg())

        if senderr != nil {
            failed = senderr
            break
        }
        log.Printf("Processed %s Successfully!", msgID)
    }
    return msgID, failed
}

func main() {
    lambda.Start(Handle)
}