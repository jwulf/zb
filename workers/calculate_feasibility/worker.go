package main

import (
	"log"
	"math/rand"

	"github.com/jwulf/zb/taskworker"
	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
)


func main() {
	taskworker.CreateWorker("calculate_feasibility", calculateFeasibility)
}

func calculateFeasibility(client worker.JobClient, job entities.Job) {
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

	payload["isFeasible"] = rand.Intn(100) > 50
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).PayloadFromMap(payload)
	if err != nil {
		// failed to set the updated payload
		taskworker.FailJob(client, job)
		return
	}

	appId := payload["appId"]
	log.Println("[", job.Type, "]", appId, " is possible: ", payload["isFeasible"])

	request.Send()
}
