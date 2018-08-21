package main

import (
    "fmt"
    //"encoding/json"


    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

//TODO: Function to create JSON schema for populating database
//TODO: Function to execute query to load schema into database 

var sess = session.Must(session.NewSessionWithOptions(session.Options{
	   	SharedConfigState: session.SharedConfigEnable,
	   }))

type eniData struct {
	pNetInt           string
	sNetInt           string
	availabilityZone  string
	eniVPC            string
	eniSubnet         string
	eniPrivateAddress string
	eniPrivateDNS     string
}

type insData struct {
	instanceType string
	instanceID   string
	tagData      []map[string]string
}

func retrieveEC2Data(instanceID ...string) ([]byte, error) { //(*ec2.DescribeInstancesOutput, error) {
	var result *ec2.DescribeInstancesOutput
	var err error
	var ec2Svc = ec2.New(sess)

	for _, id := range instanceID {
		input := &ec2.DescribeInstancesInput{
			InstanceIds: []*string{
				aws.String(id),
			},
		}	
		result, err = ec2Svc.DescribeInstances(input)
	}
	stresult := result.GoString()

	fmt.Println(stresult, err)
	return []byte(stresult), nil
}


func retrieveENIData(networkInterfaceID ...string) ([]byte, error) {//(*ec2.DescribeNetworkInterfacesOutput, error) {
	var result *ec2.DescribeNetworkInterfacesOutput
	var err error
	var ec2Svc = ec2.New(sess)

	for _, netint := range networkInterfaceID {
		input := &ec2.DescribeNetworkInterfacesInput{
			NetworkInterfaceIds: []*string{
				aws.String(netint),
			},
		}	
		
		result, err = ec2Svc.DescribeNetworkInterfaces(input)

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}	
			}	
			return nil, err
		}
	}
	
	stresult := []byte(result.GoString())
	return stresult, nil
}


func main(){
	netresult, _ := retrieveENIData("eni-e31f0f84")
	fmt.Println(string(netresult))
	//insresult, _ := retrieveEC2Data("i-0000b5f9b813c2e2a")
	//fmt.Println(string(insresult))
}

