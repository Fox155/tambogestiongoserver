package interfaces

// IHandlerDB interfaz de un handler para bases de datos
type IHandlerDB interface {
	Alta(objeto interface{}, tabla string, busqueda map[string]interface{}) error
}
