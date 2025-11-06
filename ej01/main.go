package main

import (
    "fmt"
    "pspan01/paquetes/personaje"
)

func main() {
    fmt.Println("Información del personaje:")
    fmt.Println(personaje.Descripcion())

    // Si descomentas la siguiente línea, da error de compilación:
     fmt.Println(personaje.habilidadSecreta())
}
