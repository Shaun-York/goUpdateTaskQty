package sqssrv

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// CompletionsToNetSuite stiff
type CompletionsToNetSuite struct {
    srv *sqs.SQS
}

// GetSrv return sqs
func (c *CompletionsToNetSuite) GetSrv() error {
    region := "us-east-1" //:= os.Getenv("AWS_REGION")
    awsSession, sesserr := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )
    if sesserr != nil {
        return sesserr
    }
    c.srv = sqs.New(awsSession)
    return nil
}

// Send send sqs message and remove from queue
func (c *CompletionsToNetSuite) Send(msg *sqs.SendMessageInput, delmsg *sqs.DeleteMessageInput) error {
    _, serr := c.srv.SendMessage(msg)
    if (serr != nil) {
        log.Printf("failed to send sqs message %s", *msg.MessageBody)
        return serr
    }
    _, delerr := c.srv.DeleteMessage(delmsg)
    if (delerr != nil) {
        log.Println("failed to delete sqs message")
        return delerr
    }
    return nil
}