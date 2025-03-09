package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DataDevice struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Value      float32
	PreValue   float32
	DeviceID   uuid.UUID `gorm:"type:uuid"`
	Device     Device    `gorm:"foreignKey:DeviceID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	RegisterID uuid.UUID `gorm:"type:uuid"`
	Register   Register  `gorm:"foreignKey:RegisterID"`
}

func (d *DataDevice) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.New()
	return nil
}
