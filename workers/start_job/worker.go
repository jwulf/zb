package main

import (
	"github.com/jwulf/zb-example/taskworker"
	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
	"log"
)

const BrokerAddr = "0.0.0.0:26500"

func main() {
	taskworker.CreateWorker(BrokerAddr, "start_job", startJob)
}

func startJob(client worker.JobClient, job entities.Job) {
	jobKey := job.GetKey()

	//headers, err := job.GetCustomHeadersAsMap()
	//if err != nil {
	//	// failed to handle job as we require the custom job headers
	//	failJob(client, job)
	//	return
	//}

	payload, err := job.GetPayloadAsMap()
	if err != nil {
		// failed to handle job as we require the payload
		taskworker.FailJob(client, job)
		return
	}

	payload["started"] = true
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).PayloadFromMap(payload)
	if err != nil {
		// failed to set the updated payload
		taskworker.FailJob(client, job)
		return
	}

	jobId := payload["jobId"]
	log.Println("[", job.Type, "]", jobId, " started.")

	request.Send()
}