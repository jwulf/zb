package main

import (
	"fmt"

	"github.com/jwulf/zb/broker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
)

const brokerAddr = "0.0.0.0:26500"

func main() {
	broker.EchoInfo(brokerAddr)
	//createWorkflowInstance()
	deploy()
}

func deploy() {
	zbClient, err := zbc.NewZBClient(brokerAddr)
	if err != nil {
		panic(err)
	}

	response, err := zbClient.NewDeployWorkflowCommand().AddResourceFile("howndjob.bpmn").Send()
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())
}
