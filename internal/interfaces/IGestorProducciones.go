package interfaces

import (
	"tgs/internal/models"
)

// IGestorProducciones interfaz de un gestor de producciones
type IGestorProducciones interface {
	Alta(produccion *models.Producciones) error
}
