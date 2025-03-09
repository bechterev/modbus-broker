package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Protocol struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string
}

func (p *Protocol) CreateBefore(tx gorm.DB) (err error) {
	p.ID = uuid.New()
	return nil
}
