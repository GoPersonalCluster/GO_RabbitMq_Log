package models

type ErrorLog struct {
	ID          uint   `gorm:"primaryKey"`
	Event       string `gorm:"size:64;not null"`
	Description string `gorm:"size:2048"`
	CreatedAt   int64  `gorm:"autoCreateTime"`
}
