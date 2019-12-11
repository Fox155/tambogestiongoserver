package models

import (
	"errors"

	"github.com/gookit/validate"
)

// Mensaje estructura modelo de una produccion
type Mensaje struct {
	Tambo     string `json:"Tambo" gorm:"-" validate:"required|minLen:1"`
	Sucursal  string `json:"Sucursal" gorm:"-" validate:"required|minLen:1"`
	Contenido []byte `json:"Contenido" gorm:"-" validate:"required|minLen:1"`
}

// Validacion compruebe la validez de una produccion
func (prod *Mensaje) Validacion() error {
	v := validate.Struct(prod)
	if !v.Validate() {
		return errors.New("Error en la validacion")
	}
	return nil
}
