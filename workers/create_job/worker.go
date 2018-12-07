package main

import (
	"github.com/jwulf/zb-example/taskworker"
	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
	"log"
)
const BrokerAddr = "0.0.0.0:26500"

func main() {
	taskworker.CreateWorker(BrokerAddr, "create_job", createJob)
}

func createJob(client worker.JobClient, job entities.Job) {
	jobKey := job.GetKey()
	payload, err := job.GetPayloadAsMap()
	if err != nil {
		// failed to handle job as we require the payload
		taskworker.FailJob(client, job)
		return
	}
	appId := payload["appId"]
	payload["jobId"] = appId

	log.Println("[", job.Type, "] ", appId, " created Job: ",  payload["jobId"])

	request, err := client.NewCompleteJobCommand().JobKey(jobKey).PayloadFromMap(payload)
	request.Send()
}
