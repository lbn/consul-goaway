package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/consul/api"
)

var (
	seconds    *int    = flag.Int("intervalS", 60, "Force leave failed nodes at this interval in seconds")
	consulAddr *string = flag.String("consulAddr", "", "Consul server address with port")
)

func main() {
	flag.Parse()
	client, _ := api.NewClient(&api.Config{
		Address: *consulAddr,
	})

	for {
		i := 0
		nodes, _, _ := client.Catalog().Nodes(nil)
		for _, node := range nodes {
			if !strings.Contains(node.Node, "-group-") {
				continue
			}
			hss, _, _ := client.Health().Node(node.Node, nil)
			for _, hs := range hss {
				if hs.CheckID == "serfHealth" && hs.Status == "critical" {
					client.Agent().ForceLeave(node.Node)
					fmt.Printf("Made %s leave\n", node.Node)
					i += 1
				}
			}
		}
		if i > 0 {
			fmt.Printf("Ejected %d nodes in this run\n", i)
		} else {
			fmt.Println("No failed nodes to eject in this run")
		}
		<-time.After(time.Duration(*seconds) * time.Second)
	}
}
