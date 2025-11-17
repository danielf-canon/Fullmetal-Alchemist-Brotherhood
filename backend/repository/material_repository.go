package repository

import (
	"backend-alquimia/models"
	"errors"

	"gorm.io/gorm"
)

type MaterialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) *MaterialRepository {
	return &MaterialRepository{db: db}
}

func (r *MaterialRepository) FindAll() ([]*models.Material, error) {
	var materials []*models.Material
	err := r.db.Find(&materials).Error
	if err != nil {
		return nil, err
	}
	return materials, nil
}

func (r *MaterialRepository) FindById(id int) (*models.Material, error) {
	var material models.Material
	err := r.db.Where("id = ?", id).First(&material).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &material, nil
}

func (r *MaterialRepository) Save(material *models.Material) (*models.Material, error) {
	err := r.db.Save(material).Error
	if err != nil {
		return nil, err
	}
	return material, nil
}

func (r *MaterialRepository) Delete(material *models.Material) error {
	err := r.db.Delete(material).Error
	if err != nil {
		return err
	}
	return nil
}
