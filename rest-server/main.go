package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jwulf/zb/broker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
)

const brokerAddress = "0.0.0.0:26500"

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	broker.EchoInfo(brokerAddress)
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8086", nil); err != nil {
		panic(err)
	}
}

func startWorkflow(workflowName string, payload map[string]interface{}) {
	client, err := zbc.NewZBClient(brokerAddress)
	if err != nil {
		log.Println(err)
		return
	}

	//payload := make(map[string]interface{})
	//payload["appId"] = appId

	request, err := client.NewCreateInstanceCommand().BPMNProcessId(workflowName).LatestVersion().PayloadFromMap(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg, err := request.Send()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg.String())
	}
}
