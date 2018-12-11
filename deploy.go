package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jwulf/zb/broker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
)


func main() {
	brokerPtr := flag.String("broker", "0.0.0.0", "broker address")
	workflowPtr := flag.String("workflow", "test.bpmn", "workflow file")
	flag.Parse()

	brokerAddr := *brokerPtr + ":26500"
	workflow := *workflowPtr
	log.Println("Deploying", workflow, "to", brokerAddr)

	broker.EchoInfo(brokerAddr)
	deployWorkflow(brokerAddr, workflow)
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
