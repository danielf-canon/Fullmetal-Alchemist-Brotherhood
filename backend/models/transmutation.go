package models

import (
	"backend-alquimia/api"

	"gorm.io/gorm"
)

type Transmutation struct {
	gorm.Model
	AlquimistaID uint
	Alquimista   *Alquimista  `gorm:"foreignKey:AlquimistaID;references:ID"`
	MaterialID   uint
	Material     *Material    `gorm:"foreignKey:MaterialID;references:ID"`
	Costo     	 int
	Resultado    string
	Estado       string
}

func (t *Transmutation) ToTransmutationResponseDto() *api.TransmutationResponseDto {
	return &api.TransmutationResponseDto{
		ID:            int(t.ID),
		AlquimistaID:  t.AlquimistaID,
		MaterialID:    t.MaterialID,
		Costo:         t.Costo,
		Resultado:     t.Resultado,
		Estado:        t.Estado,
		FechaCreacion: t.CreatedAt.String(),
	}
}