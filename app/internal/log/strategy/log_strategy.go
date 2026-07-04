package strategy

import (
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/config"
)

type LogStrategy struct {
	event *consumer.IntegrationEvent
}

func (pQS *LogStrategy) New(iE *consumer.IntegrationEvent) (consumer.StrategyHandler, error) {
	iE.EventName = "PII"
	mh := iE.CreateMetaHeader(config.GetHostName(), "ErrorMatchingEvent")
	mh.Args = append(mh.Args, mh.CreateArgs("NextQueue", "PII_Queue"))
	iE.MetaHeader = append(iE.MetaHeader, mh)

	return &LogStrategy{event: iE}, nil
}

func (pQS *LogStrategy) Start() ([]byte, error) {

	return []byte("PII_Queue"), nil
}
