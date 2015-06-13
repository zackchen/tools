package main

import (
	"flag"
	"fmt"
	"github.com/bitly/go-nsq"
	"log"
)

type Handler struct {
}

func (this Handler) HandleMessage(message *nsq.Message) error {

	fmt.Println("Get Response: ", string(message.Body))

	return nil

}

var (
	nsqdAddress = flag.String("nsqdAddress", "", "Nsqd address")
	topic       = flag.String("topic", "", "Topic to sub")
)

func main() {

	flag.Parse()

	channel := "123456789"
	MQConfig := nsq.NewConfig()

	consumer, err := nsq.NewConsumer(*topic, channel, MQConfig)
	if err != nil {
		log.Println("Start Consumer failed: ", err)
		return
	}
	consumer.AddHandler(Handler{})

	err = consumer.ConnectToNSQD(*nsqdAddress)
	if err != nil {
		log.Println("Error connecting nsqd: %s", err.Error())
	}
	<-consumer.StopChan
}
