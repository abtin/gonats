package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./subscribe <subject>")
		os.Exit(1)
	}
	subject := os.Args[1]
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = nc.Subscribe(subject, listener)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("press ^C to break listening for messages")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTERM, os.Interrupt)
	<-sc
	nc.Close()

}

func listener(m *nats.Msg) {
	fmt.Printf("Received: %s\n", string(m.Data))
}
