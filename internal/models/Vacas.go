package models

// Vacas estructura modelo de una vaca
type Vacas struct {
	IDVaca       int    `gorm:"column:IdVaca;primary_key"`
	IDRFID       int    `gorm:"column:IdRFID"`
	Estado       string `gorm:"column:Estado"`
	NroLactancia int    `gorm:"column:NroLactancia"`
	IDSucursal   int    `gorm:"column:IdSucursal"`
}

// Lactancias estructura modelo de una lactancia
type Lactancias struct {
	IDVaca       int `gorm:"column:IdVaca;primary_key"`
	NroLactancia int `gorm:"column:NroLactancia;primary_key"`
}
