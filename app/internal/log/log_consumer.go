package log

import (
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/log/strategy"
)

type FilterFactory struct {
	event *consumer.IntegrationEvent
}

func (c *FilterFactory) CreateStrategy(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	switch event.EventName {
	case "ErrorLog":
		return c.GetErrorQueue(event)
	case "PipelineLog":
		return c.GetPipelineQueue(event)
	}
}

func (c *FilterFactory) GetErrorQueue(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	strategy := strategy.ErrorLogStrategy{}
	return strategy.New(event)
}

func (c *FilterFactory) GetPipelineQueue(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	strategy := strategy.PipelineLogStrategy{}
	return strategy.New(event)
}
