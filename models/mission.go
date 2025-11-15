package models

import (
	"backend-alquimia/api"

	"gorm.io/gorm"
)

type Mission struct {
	gorm.Model
	Title       string 
	Description string
	Status      string     
	AssignedTo uint   //Alquimista ID
	Alquimista  *Alquimista  `gorm:"foreignKey:AssignedTo;references:ID"`
}
func (m *Mission) ToMissionResponseDto() *api.MissionResponseDto {
	return &api.MissionResponseDto{
		ID:            uint(m.ID),
		Title:         m.Title,
		Description:   m.Description,
		Status:        m.Status,
		AssignedTo:    uint(m.AssignedTo),
		CreatedAt: m.CreatedAt.String(),
	}
}