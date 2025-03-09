package repositories

import (
	"modbus-server/db/models"

	"gorm.io/gorm"
)

func CreateProtocol(db *gorm.DB, protocol *models.Protocol) error {
	return db.Create(protocol).Error
}

func DeleteProtocol(db *gorm.DB, protocol *models.Protocol) error {
	return db.Delete(protocol).Error
}
