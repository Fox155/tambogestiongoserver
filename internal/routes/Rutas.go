package routes

import (
	"tgs/internal/controllers"
	"tgs/internal/gestores"

	"github.com/gin-gonic/gin"
)

// Rutas defina las rutas que atiende el servicio
func Rutas() *gin.Engine {
	var router = gin.Default()

	gestorProducciones := gestores.GestorProducciones{2}
	controladorProducciones := controllers.ProduccionesController{Gestor: gestorProducciones}

	// Index
	router.GET("/", controladorProducciones.EstoyVivo)

	// Producciones
	router.POST("/producciones", controladorProducciones.Alta)

	return router
}
