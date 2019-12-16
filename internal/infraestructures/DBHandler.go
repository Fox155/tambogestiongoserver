package infraestructures

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBHandler estructura que contiene la informacion de la conexion actual con la base de datos
type DBHandler struct {
	Conn *gorm.DB
}

// IniciarDB inicia la configuracion con las conexiones con la base de datos y devuelve su handler
func IniciarDB() (*DBHandler, error) {
	// root:kPCR5BL3LT@/TamboGestion?charset=utf8
	options := ""
	options = options + "root"
	options = options + ":kPCR5BL3LT"
	options = options + "@tcp(localhost)"
	options = options + "/TamboGestion"
	options = options + "?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", options)

	if err != nil {
		return nil, err
	}

	return &DBHandler{db}, nil
}

// Alta permite insertar una nueva tupla en una tabla determinada
func (db *DBHandler) Alta(objeto interface{}, tabla string) error {
	tx := db.Conn.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})

	if err2 := tx.Table(tabla).Create(objeto).Error; err2 != nil {
		tx.Rollback()
		return err2
	}
	tx.Commit()
	return nil
}

// DameAlta permite insertar una nueva tupla en una tabla determinada
func (db *DBHandler) DameAlta(objeto interface{}, tabla string, busqueda map[string]interface{}) error {
	tx := db.Conn.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})
	if err2 := tx.Table(tabla).Where(busqueda).FirstOrCreate(objeto).Error; err2 != nil {
		tx.Rollback()
		return err2
	}
	tx.Commit()
	return nil
}

// CambiarEstado permite cambiar la columna 'Estado' de una tabla indicando el proximo estado y los estados validos anteriores
func (db *DBHandler) CambiarEstado(objeto interface{}, tabla string, id int64, valores map[string]interface{}, validos []string) error {
	errf := db.Conn.Table(tabla).Where(`"Estado" IN (?)`, validos).First(objeto, id).Error
	if errf != nil {
		if errf.Error() == "record not found" {
			return errors.New("No se encuentra en estado valido")
		}
		return errf
	}

	tx := db.Conn.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})

	err2 := tx.Table(tabla).Model(objeto).UpdateColumns(valores).Error

	if err2 != nil {
		tx.Rollback()
		return err2
	}
	tx.Commit()
	return nil
}

// CambiarColumnas permite cambiar las columnas de una tabla
func (db *DBHandler) CambiarColumnas(objeto interface{}, tabla string, busqueda map[string]interface{}, valores map[string]interface{}) error {
	tx := db.Conn.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: false})

	if err := tx.Table(tabla).Where(busqueda).Model(objeto).UpdateColumns(valores).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// BuscarConQuery permite realizar la busqueda con la opcion de select, order, join y busqueda por una query explicita
func (db *DBHandler) BuscarConQuery(objetos interface{}, tabla string, busqueda string, valores []interface{}, joins string, seleccion string, orden string) error {
	if err := db.Conn.Table(tabla).Order(orden).Select(seleccion).Joins(joins).Where(busqueda, valores...).Find(objetos).Error; err != nil {
		return err
	}
	return nil
}

// Dame permite instancia un objeto desde la base de datos
func (db *DBHandler) Dame(objeto interface{}, tabla string, id int64) error {
	if err := db.Conn.Table(tabla).First(objeto, id).Error; err != nil {
		return err
	}
	return nil
}

// DameBusqueda permite instancia un objeto desde la base de datos
func (db *DBHandler) DameBusqueda(objeto interface{}, tabla string, busqueda map[string]interface{}) error {
	if err := db.Conn.Table(tabla).Where(busqueda).First(objeto).Error; err != nil {
		return err
	}
	return nil
}

// DameConQuery permite realizar la busqueda con la opcion de select, order, join y busqueda por una query explicita
func (db *DBHandler) DameConQuery(objetos interface{}, tabla string, busqueda string, valores []interface{}, joins string, seleccion string, orden string) error {
	if err := db.Conn.Table(tabla).Order(orden).Select(seleccion).Joins(joins).Where(busqueda, valores...).First(objetos).Error; err != nil {
		return err
	}
	return nil
}
