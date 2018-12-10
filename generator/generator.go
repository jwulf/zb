package main

import (
	"flag"
	"fmt"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
	"log"
	"math/rand"
	"strconv"
	"time"
)


func main() {
	roundsPtr := flag.Int("rounds", 100, "number of rounds")
	brokerPtr := flag.String("broker", "0.0.0.0", "broker address")
	flag.Parse()

	fmt.Println("broker:", *brokerPtr)

	brokerAddr := *brokerPtr + ":26500"
	log.Println("Broker:", brokerAddr)
	time.Sleep(5000 * time.Millisecond)
	totalRounds := *roundsPtr
	rand.Seed(time.Now().UnixNano())
	key := strconv.Itoa(rand.Intn(100))
	for round := 0; round < totalRounds; round ++ {
		log.Println("Round:", round + 1, "of", totalRounds)
		for i := 0; i < 1000; i++ {
			createWorkflowInstance(brokerAddr, key + "-" + strconv.Itoa(round) + "-" + strconv.Itoa(i))
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func createWorkflowInstance(brokerAddr, appId string) {
	client, err := zbc.NewZBClient(brokerAddr)
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

