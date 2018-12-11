package taskworker

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jwulf/zb/broker"
	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
)

// GetBrokerAddress - check the command-line, then the environment
func getBrokerAddress() string {
	brokerPtr := flag.String("broker", "", "broker address")
	flag.Parse()
	brokerFromCmdline := *brokerPtr
	brokerFromEnv := os.Getenv("ZEEBE_BROKER_ADDRESS")
	var brokerAddress string
	log.Println("-broker command-line switch:", brokerFromCmdline)
	log.Println("ZEEBE_BROKER_ADDRESS env var:", brokerFromEnv)
	if brokerFromCmdline == "" {
		if brokerFromEnv == "" {
			brokerAddress = "0.0.0.0:26500"
		} else {
			brokerAddress = brokerFromEnv + ":26500"
		}
	} else {
		brokerAddress = brokerFromCmdline + ":26500"
	}
	return brokerAddress
}

// CreateWorker - create a worker
func CreateWorker(taskType string, handlerFn worker.JobHandler) {
	brokerAddress := getBrokerAddress()
	log.Println("Creating", taskType, "worker for broker:", brokerAddress)

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

// FailJob - call this method to fail the task
func FailJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())
	client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send()
}
