package broker

import (
	"flag"
	"fmt"
	"github.com/zeebe-io/zeebe/clients/go/pb"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
    "log"
	"os"
	"time"
)

// GetBrokerAddress - get the broker address
func GetBrokerAddress() string {
	brokerPtr := flag.String("broker", "", "broker address")
	flag.Parse()
	brokerFromCmdline := *brokerPtr
	brokerFromEnv := os.Getenv("ZEEBE_BROKER_ADDRESS")
	var brokerAddress string

	if brokerFromCmdline == "" {
		if brokerFromEnv == "" {
			brokerAddress = "0.0.0.0:26500"
		} else {
			brokerAddress = brokerFromEnv + ":26500"
		}
	} else {
		brokerAddress = brokerFromCmdline + ":26500"
	}
	log.Println("-broker command-line switch:", brokerFromCmdline)
	log.Println("ZEEBE_BROKER_ADDRESS env var:", brokerFromEnv)
	return brokerAddress
}

func EchoInfo(brokerAddress string) {
	connected := false
    for !connected {
        log.Println("Contacting Broker", brokerAddress)

        zbClient, err := zbc.NewZBClient(brokerAddress)
        if err != nil {
            log.Println(err)
            log.Println("Sleeping 5s...")
            time.Sleep(5000 * time.Millisecond)
        } else {

            topology, err := zbClient.NewTopologyCommand().Send()
            if err != nil {
                log.Println(err)
                log.Println("Sleeping 5s..")
                time.Sleep(5000 * time.Millisecond)
            } else {

                for _, broker := range topology.Brokers {
                    fmt.Println("Broker", broker.Host, ":", broker.Port)
                    for _, partition := range broker.Partitions {
                        fmt.Println("  Partition", partition.PartitionId, ":", roleToString(partition.Role))
                    }
                }
                connected = true
            }
        }
    }
}

func roleToString(role pb.Partition_PartitionBrokerRole) string {
	switch role {
	case pb.Partition_LEADER:
		return "Leader"
	case pb.Partition_FOLLOWER:
		return "Follower"
	default:
		return "Unknown"
	}
}