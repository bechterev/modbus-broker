package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DataDeviceHistory struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Value      float32
	PrevValue  float32
	DeviceID   uuid.UUID `gorm:"type:uuid;index:idx_device_register,priority:1"`
	Device     Device    `gorm:"foreignKey:DeviceID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time `gorm:"index:idx_time"`
	RegisterID uuid.UUID `gorm:"type:uuidindex:idx_device_register,priority:2"`
	Register   Register  `gorm:"foreignKey:RegisterID"`
}

func (d *DataDeviceHistory) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.New()
	return nil
}
