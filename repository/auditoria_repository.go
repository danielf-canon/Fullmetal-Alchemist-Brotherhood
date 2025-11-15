package repository

import (
	"backend-alquimia/models"
	"errors"

	"gorm.io/gorm"
)

type AuditoriaRepository struct {
	db *gorm.DB
}

func NewAuditoriaRepository(db *gorm.DB) *AuditoriaRepository {
	return &AuditoriaRepository{db: db}
}

func (r *AuditoriaRepository) FindAll() ([]*models.Auditoria, error) {
	var auditorias []*models.Auditoria
	err := r.db.Order("created_at DESC").Find(&auditorias).Error
	if err != nil {
		return nil, err
	}
	return auditorias, nil
}

func (r *AuditoriaRepository) FindById(id int) (*models.Auditoria, error) {
	var auditoria models.Auditoria
	err := r.db.Where("id = ?", id).First(&auditoria).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &auditoria, nil
}

func (r *AuditoriaRepository) Save(a *models.Auditoria) (*models.Auditoria, error) {
	err := r.db.Create(a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Nota: Normalmente las auditorías NO se eliminan.
// Lo dejamos por compatibilidad, pero puedes omitirlo si deseas.
func (r *AuditoriaRepository) Delete(a *models.Auditoria) error {
	return r.db.Delete(a).Error
}
