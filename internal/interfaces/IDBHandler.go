package interfaces

// IDBHandler interfaz de un handler para bases de datos
type IDBHandler interface {
	DameAlta(objeto interface{}, tabla string, busqueda map[string]interface{}) error
	DameConQuery(objetos interface{}, tabla string, busqueda string, valores []interface{}, joins string, seleccion string, orden string) error
	DameBusqueda(objeto interface{}, tabla string, busqueda map[string]interface{}) error
}
