package main

import (
	"fmt"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const BrokerAddr = "0.0.0.0:26500"

func main() {
	totalRounds := 100
	rand.Seed(time.Now().UnixNano())
	key := strconv.Itoa(rand.Intn(100))
	for round := 0; round < totalRounds; round ++ {
		log.Println("Round:", round + 1, "of", totalRounds)
		for i := 0; i < 1000; i++ {
			createWorkflowInstance(key + "-" + strconv.Itoa(round) + "-" + strconv.Itoa(i))
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func createWorkflowInstance(appId string) {
	client, err := zbc.NewZBClient(BrokerAddr)
	if err != nil {
		panic(err)
	}

	// After the workflow is deployed.
	payload := make(map[string]interface{})
	payload["appId"] = appId

	request, err := client.NewCreateInstanceCommand().BPMNProcessId("hownd_job").LatestVersion().PayloadFromMap(payload)
	if err != nil {
		fmt.Println(err)
	} else {

		msg, err := request.Send()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(appId, msg.String())
		}
	}
}

