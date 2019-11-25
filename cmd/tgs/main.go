package main

import (
	"fmt"
	"tgs/internal/routes"
)

func main() {
	fmt.Println("Que parece")

	r := routes.Rutas()
	r.Run(":8080")
}
