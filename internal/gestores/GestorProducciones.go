package gestores

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"tgs/internal/interfaces"
	"tgs/internal/models"
)

var once sync.Once

const (
	secret            string = "ThisIsMySuperSecret"
	tablaProducciones string = "Producciones"
	tablaSesiones     string = "SesionesOrdeño"
	tablaSucursales   string = "Sucursales"
	tablaTambos       string = "Tambos"
	tablaVacas        string = "Vacas"
	tablaEstadosVaca  string = "EstadosVacas"
	tablaLactancias   string = "Lactancias"
	tablaVacasLote    string = "VacasLote"
	tablaLotes        string = "Lotes"
)

// GestorProducciones Gestor de Producciones
type GestorProducciones struct {
	Db interfaces.IDBHandler
}

var (
	instance *GestorProducciones
)

// GestorProduccionesConstruct instancia el gestor de producciones
func GestorProduccionesConstruct() *GestorProducciones {
	once.Do(func() {
		instance = &GestorProducciones{}
	})

	return instance
}

// SetDB setea la conexion a la db
func (gestor *GestorProducciones) SetDB(db interfaces.IDBHandler) {
	gestor.Db = db
}

// Alta permite dar de alta una nueva produccion
func (gestor *GestorProducciones) Alta(tambo string, sucursalNombre string, produccion *models.Producciones) error {
	// Busco la sucursal
	sucursal := models.Sucursales{}
	var valorSucursal []interface{}
	whereS := tablaSucursales + ".Nombre = ? AND " + tablaTambos + ".Nombre = ?"
	joinS := "INNER JOIN " + tablaTambos + " USING (IdTambo)"
	seleccionS := tablaSucursales + `.*`
	valorSucursal = append(valorSucursal, sucursalNombre)
	valorSucursal = append(valorSucursal, tambo)
	if err := gestor.Db.DameConQuery(&sucursal, tablaSucursales, whereS, valorSucursal, joinS, seleccionS, ""); err != nil {
		if err.Error() == "record not found" {
			return errors.New("record not found - SUCURSAL")
		}
		return err
	}

	// Busco la vaca
	vaca := models.Vacas{}
	var valores []interface{}
	where := tablaVacas + ".IdRFID = ?" +
		" AND " + tablaEstadosVaca + ".Estado !='Muerta' AND " + tablaEstadosVaca + ".Estado !='Vendida' " +
		" AND " + tablaLactancias + ".FechaFin IS NULL" +
		" AND " + tablaVacasLote + ".FechaEgreso IS NULL"
	valores = append(valores, produccion.IDRFID)
	join := "INNER JOIN " + tablaEstadosVaca + " USING (IdVaca)" +
		" INNER JOIN " + tablaLactancias + " USING (IdVaca) " +
		" INNER JOIN " + tablaVacasLote + " USING (IdVaca) " +
		" INNER JOIN " + tablaLotes + " USING (IdLote) "
	seleccion := tablaVacas + `.*, ` + tablaLactancias + ".NroLactancia, " + tablaLotes + ".IdSucursal"
	if err := gestor.Db.DameConQuery(&vaca, tablaVacas, where, valores, join, seleccion, ""); err != nil {
		if err.Error() == "record not found" {
			return errors.New("record not found - VACA")
		}
		return err
	}

	// Busco/doyAlta la sesion de ordeño
	sesion := models.SesionesOrdeño{}
	sesion.IDSucursal = sucursal.IDSucursal
	// tiempito, errPT := time.Parse(time.RFC3339, produccion.FechaInicio.Format("2006-01-02T"))
	// if errPT != nil {
	// 	return errPT
	// }
	busquedaS := map[string]interface{}{
		"IdSucursal": sucursal.IDSucursal,
		"Fecha":      time.Date(produccion.FechaInicio.Year(), produccion.FechaInicio.Month(), produccion.FechaInicio.Day(), 0, 0, 0, 0, produccion.FechaInicio.Location()), //produccion.FechaInicio,
	}
	if err := gestor.Db.DameAlta(&sesion, tablaSesiones, busquedaS); err != nil {
		return err
	}

	// Alta de la produccion
	produccion.IDVaca = vaca.IDVaca
	produccion.NroLactancia = vaca.NroLactancia
	produccion.IDSesionOrdeño = sesion.IDSesionOrdeño
	b, errJ := json.Marshal(produccion.Medidor)
	if errJ != nil {
		return errJ
	}
	produccion.MedidorDB = b

	busqueda := map[string]interface{}{
		"IdSesionOrdeño": sesion.IDSesionOrdeño,
		"IdVaca":         vaca.IDVaca,
		"NroLactancia":   vaca.NroLactancia,
		// "Medidor":        produccion.MedidorDB,
		// "Produccion":  produccion.Produccion,
		"FechaInicio": produccion.FechaInicio,
		"FechaFin":    produccion.FechaFin,
	}

	if err := gestor.Db.DameAlta(&produccion, tablaProducciones, busqueda); err != nil {
		return err
	}

	return nil
}
