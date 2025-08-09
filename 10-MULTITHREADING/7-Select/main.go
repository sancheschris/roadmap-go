package main

import (
	"fmt"
	"sync/atomic"
	"time"
)


type Message struct {
	ID int64
	Message string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0 // the i variable can have a concurrency issue. b/c it's being shared on two threads

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 1)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	// for i := 0; i < 3; i++ {
	for {
		select {
		case msg := <-c1:
			 fmt.Printf("Received from RabbitMQ: id %d - %s\n", msg.ID, msg.Message)
		case msg := <- c2:
			fmt.Printf("Received from Kafka: id %d - %s\n", msg.ID, msg.Message)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout")
		}
	}
}
