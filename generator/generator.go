package main

import (
	"flag"
	"github.com/google/uuid"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type generator struct {
	client zbc.ZBClient
	uuid string
}

func main() {
	numJobsPtr := flag.Int("number", 10000, "number of workflow instances to start")
	brokerPtr := flag.String("brokers", "", "comma-separated broker addresses")
	workflowPtr := flag.String("workflow", "test", "workflow id")
	flag.Parse()

	brokersFromEnv := os.Getenv("ZEEBE_BROKERS")

	var brokers []string
	if *brokerPtr != "" {
		brokers = strings.Split(*brokerPtr, ",")
	} else if brokersFromEnv != "" {
		brokers = strings.Split(brokersFromEnv, ",")
	} else {
		brokers[0] = "0.0.0.0"
	}

	log.Println("Brokers:", brokers)
	log.Println("Number of jobs to start:", *numJobsPtr)
	log.Println("Workflow id:", *workflowPtr)

	workflow := *workflowPtr
	n := *numJobsPtr

	generators := make([]generator, 0)
	rand.Seed(time.Now().UnixNano())

	for _, broker := range brokers {
		client, err := zbc.NewZBClient(broker + ":26500")
		if err != nil {
			panic(err)
		}
		generatorUuid, err := uuid.NewUUID()
		var generatorId string
		if err != nil {
			log.Println(err)
			generatorId = strconv.Itoa(rand.Intn(100000))
		} else {
			generatorId = generatorUuid.String()
		}
		generators = append(generators, generator{ client, generatorId})
	}

	total := strconv.Itoa(n)
	for i := 0; i < n; i++ {
		for _, g := range generators {
			createWorkflowInstance(g.client, g.uuid+"-"+strconv.Itoa(i)+"/"+total, workflow)
		}
	}
}

func createWorkflowInstance(client zbc.ZBClient, appId string, workflowId string) {

	// After the workflow is deployed.
	payload := make(map[string]interface{})
	payload["appId"] = appId

	request, err := client.NewCreateInstanceCommand().BPMNProcessId(workflowId).LatestVersion().PayloadFromMap(payload)
	if err != nil {
		log.Println(err)
	} else {

		msg, err := request.Send()
		if err != nil {
			log.Println(err)
		} else {
			log.Println(appId, msg.String())
		}
	}
}

