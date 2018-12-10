package broker

import (
	"fmt"
	"github.com/zeebe-io/zeebe/clients/go/pb"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
    "log"
    "time"
)

func EchoInfo(BrokerAddr string) {
    connected := false
    for !connected {
        log.Println("Contacting Broker", BrokerAddr)

        zbClient, err := zbc.NewZBClient(BrokerAddr)
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