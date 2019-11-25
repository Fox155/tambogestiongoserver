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
