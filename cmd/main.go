package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
	"github.com/ruhancs/listen-events/internal/application/dto"
	"github.com/ruhancs/listen-events/internal/application/factory"
	elastic "github.com/ruhancs/listen-events/internal/infra/database/elasticksearch"
	rabbitmq "github.com/ruhancs/listen-events/pkg/queue"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	elkClient := elastic.ConnectWithElasticSearch(context.Background())

	registerEventUseCase := factory.RegisterEventUseCaseFactory(elkClient)
	registerLogErrorUseCase := factory.RegisterLogErrorUseCaseFactory(elkClient)
	
	ch,err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	
	msgsOutEvent := make(chan amqp091.Delivery)
	msgsOutLogError := make(chan amqp091.Delivery)

	go rabbitmq.Consumer(ch,msgsOutEvent,"events")
	go rabbitmq.Consumer(ch,msgsOutLogError,"log_errors")
	
	go func ()  {
		for msg := range msgsOutEvent {
			fmt.Println("event received")
			var inputRegisterEvent dto.RegisterEventInputDto
			err := json.Unmarshal(msg.Body,&inputRegisterEvent)
			if err != nil {
				//criar dead letter queue
				log.Panic(err)
			}
			outputMsg,err := registerEventUseCase.Execute(context.Background(),inputRegisterEvent)
			if err != nil {
				//criar dead letter queue
				log.Panic(err)
			}
			log.Println(outputMsg)
			//apagar msg da fila
			msg.Ack(false)
		}
		}()
		
	for msg := range msgsOutLogError {
		var inputRegisterLogError dto.RegisterLogErrortInputDto
		err := json.Unmarshal(msg.Body,&inputRegisterLogError)
		if err != nil {
			//criar dead letter queue
			log.Panic(err)
		}
		outputMsg,err := registerLogErrorUseCase.Execute(context.Background(),inputRegisterLogError)
		if err != nil {
			//criar dead letter queue
			log.Panic(err)
		}
		log.Println(outputMsg)
		//apagar msg da fila
		msg.Ack(false)
	}
}