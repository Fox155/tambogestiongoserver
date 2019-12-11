package models

import (
	"errors"
	"time"

	"github.com/gookit/validate"
)

// Producciones estructura modelo de una produccion
type Producciones struct {
	IDProduccion   int64             `gorm:"column:IdProduccion;primary_key"`
	IDSesionOrdeño int64             `gorm:"column:IdSesionOrdeño"`
	IDVaca         int               `gorm:"column:IdVaca"`
	NroLactancia   int               `gorm:"column:NroLactancia"`
	Produccion     float32           `gorm:"column:Produccion"`
	FechaInicio    time.Time         `gorm:"column:FechaInicio"`
	FechaFin       time.Time         `gorm:"column:FechaFin"`
	Medidor        map[string]string `gorm:"-"`
	MedidorDB      []byte            `gorm:"column:Medidor"`
	IDRFID         int               `gorm:"-"`
}

// Validacion compruebe la validez de una produccion
func (prod *Producciones) Validacion() error {
	v := validate.Struct(prod)
	if !v.Validate() {
		return errors.New("Error en la validacion")
	}
	return nil
}
