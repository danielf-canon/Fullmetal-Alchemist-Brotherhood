package models

import (
	"backend-alquimia/api"

	"gorm.io/gorm"
)

type Alquimista struct {
	gorm.Model
	
	Nombre        string
	Edad          int
	Especialidad  string
	Rango         string
}


func (a *Alquimista) ToAlquimistaResponseDto() *api.AlquimistaResponseDto {
	return &api.AlquimistaResponseDto{
		ID:            int(a.ID),
		Nombre:        a.Nombre,
		Edad:          a.Edad,
		Especialidad:  a.Especialidad,
		Rango:         a.Rango,
		FechaCreacion: a.CreatedAt.String(),
	}
}
