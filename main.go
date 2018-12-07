package main

import (
	"fmt"
	"github.com/jwulf/zb-example/broker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
)
const BrokerAddr = "0.0.0.0:26500"

func main() {
	broker.EchoInfo(BrokerAddr)
	//createWorkflowInstance()
	deploy()
}


func deploy() {
	zbClient, err := zbc.NewZBClient(BrokerAddr)
	if err != nil {
		panic(err)
	}

	response, err := zbClient.NewDeployWorkflowCommand().AddResourceFile("howndjob.bpmn").Send()
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())
}