package models

import (
	"errors"
	"time"

	"github.com/gookit/validate"
)

// Producciones estructura modelo de una produccion
type Producciones struct {
	Tambo          string `json:"Tambo" validate:"required|minLen:1"`
	Mensaje        []byte `json:"Mensaje" validate:"required|minLen:1"`
	IDProduccion   int64
	IDSesionOrdeño int64
	IDVaca         int
	NroLactancia   int
	Produccion     float32
	FechaInicio    time.Time
	FechaFin       time.Time
	Medidor        map[string]string
	MedidorDB      []byte
}

// Validacion compruebe la validez de una produccion
func (prod *Producciones) Validacion() error {
	v := validate.Struct(prod)
	if !v.Validate() {
		return errors.New("Error en la validacion")
	}
	return nil
}
