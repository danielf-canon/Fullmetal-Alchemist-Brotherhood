package repository

import (
	"backend-alquimia/models"
	"errors"

	"gorm.io/gorm"
)

type AlquimistaRepository struct {
	db *gorm.DB
}

func NewAlquimistaRepository(db *gorm.DB) *AlquimistaRepository {
	return &AlquimistaRepository{
		db: db,
	}
}

func (p *AlquimistaRepository) FindAll() ([]*models.Alquimista, error) {
	var alquimistas []*models.Alquimista
	err := p.db.Find(&alquimistas).Error
	if err != nil {
		return nil, err
	}
	return alquimistas, nil
}

func (p *AlquimistaRepository) FindById(id int) (*models.Alquimista, error) {
	var alquimista models.Alquimista
	err := p.db.Where("id = ?", id).First(&alquimista).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &alquimista, nil
}

func (p *AlquimistaRepository) Save(data *models.Alquimista) (*models.Alquimista, error) {
	err := p.db.Save(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *AlquimistaRepository) Delete(data *models.Alquimista) error {
	err := p.db.Delete(data).Error
	if err != nil {
		return err
	}
	return nil
}
