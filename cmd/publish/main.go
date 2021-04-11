package main

import (
	"bufio"
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./publish <subject> <file_to_publish")
		os.Exit(1)
	}
	subject := os.Args[1]
	publishFile := os.Args[2]

	var paused bool
	ch := make(chan bool)
	go func(p chan bool) {
		for {
			var command string
			fmt.Print("command [pause|resume]: ")
			l, err := fmt.Scanf("%s", &command)
			if l == 0 {
				continue
			}
			if err != nil {
				fmt.Println("Error reading input - %s", err)
			}
			switch strings.ToLower(command) {
			case "pause":
				paused = true
			case "resume":
				paused = false
				p <- false
			}
		}
	}(ch)

	file, err := os.Open(publishFile)
	if err != nil {
		fmt.Printf("cannot open file %q to publish\n", publishFile)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for scanner.Scan() {
		if !paused {
			line := scanner.Bytes()
			if err := nc.Publish(subject, line); err != nil {
				fmt.Printf("Error publishing message - %s\n", err)
			}
			time.Sleep(1 * time.Second) // to allow demonstrating pause/resume
		} else {
			paused = <-ch
		}
	}
}
