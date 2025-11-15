package repository

import (
	"backend-alquimia/models"
	"errors"

	"gorm.io/gorm"
)

type MissionRepository struct {
	db *gorm.DB
}

func NewMissionRepository(db *gorm.DB) *MissionRepository {
	return &MissionRepository{
		db: db,
	}
}

func (r *MissionRepository) FindAll() ([]*models.Mission, error) {
	var missions []*models.Mission
	err := r.db.Preload("Alquimista").Find(&missions).Error
	if err != nil {
		return nil, err
	}
	return missions, nil
}

func (r *MissionRepository) FindById(id int) (*models.Mission, error) {
	var mission models.Mission
	err := r.db.Preload("Alquimista").Where("id = ?", id).First(&mission).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &mission, nil
}

func (r *MissionRepository) Save(mission *models.Mission) (*models.Mission, error) {
	err := r.db.Save(mission).Error
	if err != nil {
		return nil, err
	}
	return mission, nil
}

func (r *MissionRepository) Delete(mission *models.Mission) error {
	err := r.db.Delete(mission).Error
	if err != nil {
		return err
	}
	return nil
}
