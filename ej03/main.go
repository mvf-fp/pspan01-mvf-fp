package main

import (
    "fmt"
    "pspan01/paquetes/task"
)

func main() {
    fmt.Println("Añadiendo tareas...")

    lista := task.TaskList{}

    lista.AddTask("Aprender Go")
    lista.AddTask("Hacer la compra")
    lista.AddTask("Preparar la cena")

    lista.ListTasks()

    fmt.Println("Completando la tarea 1...")
    lista.CompleteTask(1)
    lista.ListTasks()
}
