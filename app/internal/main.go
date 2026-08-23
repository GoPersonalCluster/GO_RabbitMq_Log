package main

import (
	"net/http"
	stdlog "log"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service"
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/database"
	applog "github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/log"
)

func main() {
	database.Connect()
	service := service.RabbitMQConfigComposite{}
	service.ConfigureConnection()

	logCommand := applog.LogFactory{}

	filterConsumer := consumer.FilterConsumer{}
	config := consumer.ConsumerConfig{}

	config.AbstractFactory = &logCommand
	config.Durable = false
	config.Exclusive = false
	config.AutoDelete = false
	config.NoWait = true
	config.QueueName = "log_queue"
	config.Args = nil

	filterConsumer.SetConfiguration(&config)

	service.AddConsumer("log_queue", &filterConsumer)
	service.Start()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	stdlog.Println("[main] servidor HTTP ouvindo na porta 8080")
	stdlog.Fatal(http.ListenAndServe(":8080", nil))
}