package models

import (
	"time"
)

// SesionesOrdeño estructura modelo de una sesion de ordeño
type SesionesOrdeño struct {
	IDSesionOrdeño int64     `gorm:"column:IdSesionOrdeño;primary_key"`
	IDSucursal     int       `gorm:"column:IdSucursal"`
	Fecha          time.Time `gorm:"-"`
	Observaciones  string    `gorm:"column:Observaciones"`
}

// Sucursales estructura modelo de una sucursal
type Sucursales struct {
	IDSucursal int    `gorm:"column:IdSucursal;primary_key"`
	Nombre     string `gorm:"column:Nombre"`
}
