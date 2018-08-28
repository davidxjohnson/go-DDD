package main

import (
	"fmt"

	"bitbucket/awshandlers"
	
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/satori/go.uuid"
)

func createMessage(payload, queue string) *sqs.SendMessageInput {
	var msg sqs.SendMessageInput

	msg.SetMessageBody(payload)
	msg.SetMessageDeduplicationId(uuid.Must(uuid.NewV4()).String())	
	msg.SetQueueUrl(queue)
	msg.SetMessageGroupId("1")

	return &msg
}

func main() {
	sqs := awshandlers.CreateSQSService()	

	output, err := sqs.SendMessage(createMessage(
		"helloSQS!",
		"https://sqs.us-east-1.amazonaws.com/330932647836/SERVICE-MAP-DLQ.fifo"))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(output)
}

