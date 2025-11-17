package models

import (
	"backend-alquimia/api"
	"gorm.io/gorm"
)

type Auditoria struct {
	gorm.Model
	User        string
	Accion      string
	Entidad     string
	Descripcion string
}

func (a *Auditoria) ToAuditoriaResponseDto() *api.AuditoriaResponseDto {
	return &api.AuditoriaResponseDto{
		ID:            int(a.ID),
		User:          a.User,
		Accion:        a.Accion,
		Entidad:       a.Entidad,
		Descripcion:   a.Descripcion,
		FechaCreacion: a.CreatedAt.String(),
	}
}
