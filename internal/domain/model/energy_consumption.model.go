package model

import (
	"time"

	"github.com/google/uuid"
)

type EnergyConsumption struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	DeviceId    uint8      `gorm:"not null"`
	Consumption float32   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
