package gestores

import (
	"tgs/internal/models"
)

// GestorProducciones Gestor de Producciones
type GestorProducciones struct {
	// DB interfaces.IHandlerDB
	Dato int
}

// Alta permite dar de alta una nueva produccion
func (gestor *GestorProducciones) Alta(produccion *models.Producciones) error {
	return nil
}
