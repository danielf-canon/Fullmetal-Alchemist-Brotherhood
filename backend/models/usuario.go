package models

import "gorm.io/gorm"

type Usuario struct {
    gorm.Model
    Name         string
    Email        string `gorm:"unique"`
    PasswordHash string
    Rol          string  
    AlquimistaID *uint       
    Alquimista   *Alquimista `gorm:"foreignKey:AlquimistaID;references:ID"`
}
