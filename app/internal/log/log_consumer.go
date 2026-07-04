package log

import (
	"errors"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/config"
	"github.com/GoPersonalCluster/go_rabbitmq_log/app/internal/log/strategy"
)

type FilterFactory struct {
	event *consumer.IntegrationEvent
}

func (c *FilterFactory) CreateStrategy(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {

	switch event.EventName {
	case "PII":
		return c.GetPIIQueue(event)
	default:
		return nil, c.GetDefaultErrorResponse(event)
	}
}

func (c *FilterFactory) GetDefaultErrorResponse(event *consumer.IntegrationEvent) error {
	event.CreateMetaHeader(config.GetHostName(), "ErrorMatchingEvent")
	return errors.New(event.EventName + "event not found")
}


func (c *FilterFactory) GetPIIQueue(event *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	strategy := strategy.LogStrategy{}
	

	return strategy.New(event)
}
