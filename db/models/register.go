package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Register struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string
	ProtocolID  uuid.UUID
	Protocol    Protocol `gorm:"foreignKey:ProtocolID"`
	Address     uint
	Type        byte
	SensPercent float32
	SensValue   float32
	Value       float32
}

func (r *Register) CreateBefore(tx gorm.DB) (err error) {
	r.ID = uuid.New()
	return nil
}
