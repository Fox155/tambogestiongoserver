package models

import (
	"time"
)

// SesionesOrdeño estructura modelo de una sesion de ordeño
type SesionesOrdeño struct {
	IDSesionOrdeño int64
	IDSucursal     int
	Fecha          time.Time
	Observaciones  string
}

// Sucursales estructura modelo de una sucursal
type Sucursales struct {
	IDSucursal int    `gorm:"column:IdSucursal;primary_key"`
	Nombre     string `gorm:"column:Nombre"`
}
