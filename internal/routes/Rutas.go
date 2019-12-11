package routes

import (
	"tgs/internal/controllers"
	"tgs/internal/gestores"
	"tgs/internal/interfaces"

	"github.com/gin-gonic/gin"
)

// Rutas defina las rutas que atiende el servicio
func Rutas(db interfaces.IDBHandler) *gin.Engine {
	var router = gin.Default()

	gestorProducciones := gestores.GestorProducciones{Db: db}
	controladorProducciones := controllers.ProduccionesController{Gestor: gestorProducciones}

	// Producciones
	router.POST("/producciones", controladorProducciones.Alta)

	// Index
	router.GET("/", controladorProducciones.EstoyVivo)

	return router
}
