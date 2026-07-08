package strategy

import (
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/config"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/database"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/models"
)

type PipelineLogStrategy struct {
	event *consumer.IntegrationEvent
}

func (pQS *PipelineLogStrategy) New(iE *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	iE.EventName = "PII"
	mh := iE.CreateMetaHeader(config.GetHostName(), "PipelineLogEvent")
	iE.MetaHeader = append(iE.MetaHeader, mh)

	return &PipelineLogStrategy{event: iE}, nil
}

func (pQS *PipelineLogStrategy) Start() ([]byte, error) {

	for _, header := range pQS.event.MetaHeader {

		log := models.NewPipelineLog(pQS.event.ID, header.Source, header.EventName, header.Args)
		database.DB.Create(&log)
	}

	return nil, nil
}
