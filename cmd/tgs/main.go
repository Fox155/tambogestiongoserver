package main

import (
	"fmt"
	"tgs/internal/gestores"
	"tgs/internal/infraestructures"
	"tgs/internal/routes"
)

// func main2() {
// 	fmt.Println("Que parece")

// 	r := routes.Rutas()
// 	r.Run(":8080")
// }

func main() {
	conf, err := initConfig()
	if err != nil {
		panic(err)
	}

	db, err := infraestructures.IniciarDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión a la db exitosa")

	defer db.Conn.Close()

	db.Conn.LogMode(true)
	g := gestores.GestorProduccionesConstruct()
	g.SetDB(db)

	r := routes.Rutas(db)

	errR := r.Run(conf.Server.Listen)
	if errR != nil {
		panic(errR)
	}
}
