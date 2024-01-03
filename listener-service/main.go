package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/eduardylopes/go-microservices/listener-service/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	log.Println("listening for and consuming rabbitmq messages...")

	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	err = consumer.Listen([]string{"log.INFO", "log.WARN", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (conn *amqp.Connection, err error) {
	counts := 0
	backOff := 1 * time.Second

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("rabbitmq not yet ready...")
			counts++
		} else {
			fmt.Println("connected to rabbitmq")
			conn = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
	}

	return
}
