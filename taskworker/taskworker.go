package taskworker

import (
	"fmt"
	"github.com/jwulf/zb-example/broker"
	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
    "log"
    "time"
)

func CreateWorker(brokerAddress string, taskType string, handlerFn worker.JobHandler) {
    fmt.Println("createWorker")
	broker.EchoInfo(brokerAddress)

	client := getClient(brokerAddress)

	fmt.Println("Listening for", taskType)

	createJobWorker := client.NewJobWorker().JobType(taskType).Handler(handlerFn).Open()
	defer createJobWorker.Close()
	createJobWorker.AwaitClose()
}

func getClient(brokerAddress string) zbc.ZBClient {
    connected := false
    var client zbc.ZBClient
    var err error

    for !connected {
        log.Println("Getting client...")
      client, err = zbc.NewZBClient(brokerAddress)
      if err != nil {
          log.Println(err)
          time.Sleep(1000 * time.Millisecond)
      } else {
          connected = true
      }
    }
    return client
}

func FailJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())
	client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send()
}
