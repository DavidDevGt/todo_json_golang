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
	case "eliminar": // Nuevo comando para eliminar tarea
		if len(args) < 2 {
			fmt.Println("Uso: mi_app_en_go eliminar <ID>")
			return
		}
		id := args[1]
		EliminarTarea(id)
	case "ayuda": // Comando para mostrar la ayuda
		MostrarAyuda()
	default:
		fmt.Println("Comando no reconocido, usa 'ayuda' para ver los comandos disponibles.")
	}
}

func MostrarAyuda() {
	fmt.Println("Uso de la aplicación:")
	fmt.Println("todo_json listar: Lista todas las tareas.")
	fmt.Println("todo_json agregar <tarea>: Agrega una nueva tarea.")
	fmt.Println("todo_json completar <ID>: Marca una tarea como completada.")
	fmt.Println("todo_json eliminar <ID>: Elimina una tarea.")
	fmt.Println("todo_json ayuda: Muestra esta ayuda.")
}
