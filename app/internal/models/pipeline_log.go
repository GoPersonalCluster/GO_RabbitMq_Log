package models

import (
	"encoding/json"
	"time"

	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/google/uuid"
)

type PipelineLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	EventID   string    `gorm:"size:64;not null"`
	Source    string    `gorm:"size:64;not null"`
	EventName string    `gorm:"size:64;not null"`
	Args      string    `gorm:"size:64;not null"`
	OccuredAt int64     `gorm:"autoCreateTime"`
}

func NewPipelineLog(eventID string, source string, eventName string, args []consumer.Args) PipelineLog {
	argsJSON, _ := json.Marshal(args)

	return PipelineLog{
		ID:        uuid.New(),
		EventID:   eventID,
		Source:    source,
		EventName: eventName,
		Args:      string(argsJSON),
		OccuredAt: time.Now().Unix(),
	}
}
