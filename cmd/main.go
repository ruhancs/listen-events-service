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
	"github.com/ruhancs/listen-events/internal/infra/web"
	rabbitmq "github.com/ruhancs/listen-events/pkg/queue"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	elkClient := elastic.ConnectWithElasticSearch(context.Background())

	//registerEventUseCase := factory.RegisterEventUseCaseFactory(elkClient)
	bulkRegisterEventUseCase := factory.BulkRegisterEventUseCaseFactory(elkClient)
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
		var eventsToRegister []dto.RegisterEventInputDto
		for msg := range msgsOutEvent {
				var inputRegisterEvent dto.RegisterEventInputDto
				err := json.Unmarshal(msg.Body,&inputRegisterEvent)
				if err != nil {
					//criar dead letter queue
					log.Panic(err)
				}
				//msg.Ack(false)
				eventsToRegister = append(eventsToRegister, inputRegisterEvent)
				fmt.Println(len(eventsToRegister))
				if len(eventsToRegister) == 3 {
					fmt.Println("Proccess All Events")
					outputMsg,err := bulkRegisterEventUseCase.Execute(context.Background(),eventsToRegister)
					if err != nil {
						//criar dead letter queue
						log.Panic(err)
					}
					log.Println(outputMsg)
					eventsToRegister = []dto.RegisterEventInputDto{}
				}
				//apagar msg da fila
				msg.Ack(false)
			}
		}()
	
	go func ()  {
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
	}()
	
	searchEventUseCase := factory.SearchEventUseCaseFactory(elkClient)
	searchLogErrorUseCase := factory.SearchLogErrorUseCaseFactory(elkClient)
	app := web.NewApplication(searchEventUseCase,searchLogErrorUseCase)

	app.Server()
}