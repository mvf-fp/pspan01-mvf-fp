package task

import "fmt"

// Struct público Task
type Task struct {
    ID        int
    Title     string
    Completed bool
}

// Struct público TaskList
type TaskList struct {
    Tasks []*Task
}

// Método para añadir una tarea (receptor puntero)
func (tl *TaskList) AddTask(title string) {
    newID := len(tl.Tasks) + 1
    task := &Task{
        ID:        newID,
        Title:     title,
        Completed: false,
    }
    tl.Tasks = append(tl.Tasks, task)
}

// Método para completar una tarea por ID (receptor puntero)
func (tl *TaskList) CompleteTask(id int) {
    for _, t := range tl.Tasks {
        if t.ID == id {
            t.Completed = true
            return
        }
    }
    fmt.Println("No se encontró la tarea con ID:", id)
}

// Método para listar tareas (receptor de valor)
func (tl TaskList) ListTasks() {
    fmt.Println("--- Lista de Tareas ---")
    for _, t := range tl.Tasks {
        estado := "Pendiente"
        if t.Completed {
            estado = "Completada"
        }
        fmt.Printf("ID: %d, Título: %s, Estado: %s\n", t.ID, t.Title, estado)
    }
    fmt.Println("-----------------------")
}
