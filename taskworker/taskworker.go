package taskworker

import (
	"fmt"
	"github.com/jwulf/zb-example/broker"
	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
	"log"
)

func CreateWorker(brokerAddress string, taskType string, handlerFn worker.JobHandler) {
	broker.EchoInfo(brokerAddress)

	client, err := zbc.NewZBClient(brokerAddress)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening for", taskType)

	createJobWorker := client.NewJobWorker().JobType(taskType).Handler(handlerFn).Open()
	defer createJobWorker.Close()
	createJobWorker.AwaitClose()
}

func FailJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())
	client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send()
}
