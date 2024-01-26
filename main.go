package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Uso: mi_app_en_go <comando>")
		return
	}

	comando := args[0]

	switch comando {
	case "listar":
		ListarTareas()
	case "agregar":
		if len(args) < 2 {
			fmt.Println("Uso: mi_app_en_go agregar <tarea>")
			return
		}
		titulo := args[1]
		AgregarTarea(titulo)
	case "completar":
		if len(args) < 2 {
			fmt.Println("Uso: mi_app_en_go completar <ID>")
			return
		}
		id := args[1]
		CompletarTarea(id)
	default:
		fmt.Println("Comando no reconocido")
	}
}
