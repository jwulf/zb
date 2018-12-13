package main

import (
	"flag"
	"fmt"
	"github.com/jwulf/zb/broker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
	"log"
)


func main() {
	workflowPtr := flag.String("workflow", "test.bpmn", "workflow file")
	flag.Parse()

	workflow := *workflowPtr

	brokerAddress := broker.GetBrokerAddress()

	log.Println("Deploying", workflow, "to", brokerAddress)

	broker.EchoInfo(brokerAddress)
	deployWorkflow(brokerAddress, workflow)
}

func deployWorkflow(brokerAddr string, workflow string) {
	zbClient, err := zbc.NewZBClient(brokerAddr)
	if err != nil {
		panic(err)
	}

	response, err := zbClient.NewDeployWorkflowCommand().AddResourceFile(workflow).Send()
	if err != nil {
		log.Println("An error occurred while deploying the workflow...")
		panic(err)
	}

	fmt.Println(response.String())
}
