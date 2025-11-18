package repository

import (
    "backend-alquimia/models"
    "errors"

    "gorm.io/gorm"
)

type UsuarioRepository struct {
    db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
    return &UsuarioRepository{db: db}
}

func (r *UsuarioRepository) FindByEmail(email string) (*models.Usuario, error) {
    var user models.Usuario
    err := r.db.Where("email = ?", email).First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &user, err
}

func (r *UsuarioRepository) Save(u *models.Usuario) (*models.Usuario, error) {
    err := r.db.Save(u).Error
    return u, err
}
