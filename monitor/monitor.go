package main

import (
	"github.com/jwulf/zb/broker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
	"log"
)

const brokerAddress = "nukkit.magikcraft.io:26500"
const workflowId = "hownd_job"

func main() {
	broker.EchoInfo(brokerAddress)

	client, _ := zbc.NewZBClient(brokerAddress)

	result, err := client.NewGetWorkflowCommand().BpmnProcessId(workflowId).LatestVersion().Send()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(result)
}
