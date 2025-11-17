package repository

import (
	"backend-alquimia/models"
	"errors"

	"gorm.io/gorm"
)

type TransmutationRepository struct {
	db *gorm.DB
}

func NewTransmutationRepository(db *gorm.DB) *TransmutationRepository {
	return &TransmutationRepository{
		db: db,
	}
}

func (r *TransmutationRepository) FindAll() ([]*models.Transmutation, error) {
	var transmutations []*models.Transmutation
	err := r.db.Preload("Alquimista").Preload("Material").Find(&transmutations).Error
	if err != nil {
		return nil, err
	}
	return transmutations, nil
}

func (r *TransmutationRepository) FindById(id int) (*models.Transmutation, error) {
	var transmutation models.Transmutation
	err := r.db.Preload("Alquimista").Preload("Material").Where("id = ?", id).First(&transmutation).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &transmutation, nil
}

func (r *TransmutationRepository) Save(data *models.Transmutation) (*models.Transmutation, error) {
	err := r.db.Save(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *TransmutationRepository) Delete(data *models.Transmutation) error {
	err := r.db.Delete(data).Error
	if err != nil {
		return err
	}
	return nil
}
