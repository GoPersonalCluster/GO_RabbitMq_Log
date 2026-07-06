package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PipelineLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	EventID   string `gorm:"size:64;not null"`
	Source    string `gorm:"size:64;not null"`
	EventName string `gorm:"size:64;not null"`
	Args      string `gorm:"size:64;not null"`
	OccuredAt int64  `gorm:"autoCreateTime"`
}

func NewPipelineLog() PipelineLog {



}