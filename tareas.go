package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Tarea struct {
	ID         int    `json:"id"`
	Titulo     string `json:"titulo"`
	Completada bool   `json:"completada"`
}

var tareas []Tarea

func CargarTareas() {
	data, err := ioutil.ReadFile("tareas.json")
	if err != nil {
		fmt.Println("Error al cargar las tareas:", err)
		return
	}

	err = json.Unmarshal(data, &tareas)
	if err != nil {
		fmt.Println("Error al decodificar el archivo JSON:", err)
		return
	}
}

func GuardarTareas() {
	data, err := json.Marshal(tareas)
	if err != nil {
		fmt.Println("Error al codificar las tareas:", err)
		return
	}

	err = ioutil.WriteFile("tareas.json", data, 0644)
	if err != nil {
		fmt.Println("Error al guardar las tareas:", err)
	}
}

func ListarTareas() {
	CargarTareas()
	for _, tarea := range tareas {
		fmt.Printf("ID: %d - Título: %s - Completada: %v\n", tarea.ID, tarea.Titulo, tarea.Completada)
	}
}

func AgregarTarea(titulo string) {
	CargarTareas()
	nuevaTarea := Tarea{
		ID:         len(tareas) + 1,
		Titulo:     titulo,
		Completada: false,
	}
	tareas = append(tareas, nuevaTarea)
	GuardarTareas()
	fmt.Println("Tarea agregada con éxito.")
}

func CompletarTarea(id string) {
	CargarTareas()
	for i, tarea := range tareas {
		if fmt.Sprintf("%d", tarea.ID) == id {
			tareas[i].Completada = true
			GuardarTareas()
			fmt.Println("Tarea marcada como completada.")
			return
		}
	}
	fmt.Println("No se encontró la tarea con el ID especificado.")
}
