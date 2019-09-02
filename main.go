package main

import (
	"log"
	"os"
	//"runtime"
	//"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

const (
	topic = "discount"
)

func main() {

	log.Println("... Subscribing to topic:", topic)

	natsServer := os.Getenv("ECOMMERCE_NATS_SERVICE_HOST")
	log.Println("NAT server --- " + natsServer)

	natsURL := "nats://" + natsServer + ":4222"
	log.Println("NAT server conn URL --- " + natsURL)

	// opts := nats.Options{
	// 	AllowReconnect: true,
	// 	MaxReconnect:   5,
	// 	ReconnectWait:  5 * time.Second,
	// 	Timeout:        3 * time.Second,
	// 	Url:            natsURL,
	// }

	//opts := []nats.Option{nats.Name("NATS ecommerce subscriber")}
	// conn, _ := opts.Connect()
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	// keep it open 
	// defer nc.Close()

	// ecommerce-stan is the cluster name
	// $ kubectl get crd
	// $ kubectl get natsstreamingclusters.streaming.nats.io 
	sc, err := stan.Connect("ecommerce-stan", "sub-sales-discount-1", stan.NatsConn(nc))
	if err != nil {
		log.Fatal(err)
	}
	// Keep it open
	//defer sc.Close()

	//log.Println("Subscriber connected to NATS server")

	//log.Printf("Subscribing to topic %s\n", topic)
	sc.Subscribe(topic, func(msg *stan.Msg) {
		log.Printf("Topic: %s received message: %s\n", topic, string(msg.Data))
	}, stan.DeliverAllAvailable())

	select{}
}
