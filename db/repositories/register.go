package repositories

import (
	"modbus-server/db/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateRegister(db *gorm.DB, register *models.Register) error {
	return db.Create(register).Error
}

func UpdateRegister(db *gorm.DB, register *models.Register) error {
	return db.Save(register).Error
}

func GetRegister(db *gorm.DB, id uuid.UUID) (models.Register, error) {
	var register models.Register
	err := db.First(&register, "id=?", id).Error
	if err == nil {
		return models.Register{}, nil
	}
	return register, err
}

func DeleteRegister(db *gorm.DB, register *models.Register) error {
	return db.Delete(register).Error
}
