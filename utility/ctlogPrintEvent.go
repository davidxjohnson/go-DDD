package main

import (
	"fmt"
	"encoding/base64"
	"bytes"

	"compress/gzip"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/lambda"
)

type receivedEvent struct {
	AwsLogs eventData `json:"awslogs"`
}

type eventData struct {
	Data string  `json:"data"`
}

func decompressEvent(data eventData) []byte {
	decodedEvent, _ := base64.StdEncoding.DecodeString(data.Data) 
	inputStream := bytes.NewBuffer(decodedEvent)

	r, err := gzip.NewReader(inputStream)
	defer r.Close()

	if err != nil {
		fmt.Println(err)
	}

	decompressedEvent, err := ioutil.ReadAll(r)

	return decompressedEvent
}

func printEvent(event receivedEvent) {
		fmt.Println(decompressEvent(event.AwsLogs))
}

func main() {
	lambda.Start(printEvent)
}
