package models

import (
	"backend-alquimia/api"

	"gorm.io/gorm"
)

type Material struct {
	gorm.Model
	NombreMaterial         string 
	Transmutations []Transmutation `gorm:"foreignKey:MaterialID"`
}


func (m *Material) ToMaterialResponseDto() *api.MaterialResponseDto {
	return &api.MaterialResponseDto{
		ID:             int(m.ID),
		NombreMaterial: m.NombreMaterial,
		FechaCreacion:  m.CreatedAt.String(),
	}
}
