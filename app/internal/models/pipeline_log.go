package models

type PipelineLog struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"size:2048"`
	CreatedAt   int64  `gorm:"autoCreateTime"`
}
