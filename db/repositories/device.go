package repositories

import (
	"modbus-server/db/models"

	"gorm.io/gorm"
)

func GetAllDevices(db *gorm.DB) []models.Device {
	var devices []models.Device
	db.Find(&devices)
	return devices
}

func CreateDevice(db *gorm.DB, device *models.Device) error {
	return db.Create(device).Error
}

func UpdateDevice(db *gorm.DB, device *models.Device) error {
	return db.Save(device).Error
}

func DeleteDevice(db *gorm.DB, device *models.Device) error {
	return db.Delete(device).Error
}
