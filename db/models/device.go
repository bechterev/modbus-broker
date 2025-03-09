package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name          string
	Address       int
	Ip            string
	BadConnection bool
	ProtocolID    uuid.UUID `gorm:"type:uuid"`
	Protocol      Protocol  `gorm:"foreignKey:ProtocolID"`
	Active        bool
}

func (d *Device) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.New()
	return nil
}
