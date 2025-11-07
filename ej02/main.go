package main

import (
    "fmt"
    "pspan01/paquetes/salario"
)

func main() {
    fmt.Println("--- Caso 1: Éxito ---")
    if salarioCalc, err := salario.CalcularSalario(100, 20); err != nil {
        fmt.Println("Ocurrió un error:", err)
    } else {
        fmt.Printf("El salario calculado es: %.2f€\n\n", salarioCalc)
    }

    fmt.Println("--- Caso 2: Error por horas insuficientes ---")
    if salarioCalc, err := salario.CalcularSalario(70, 20); err != nil {
        fmt.Println("Ocurrió un error:", err)
    } else {
        fmt.Printf("El salario calculado es: %.2f€\n\n", salarioCalc)
    }

    fmt.Println("--- Caso 3: Error por coste bajo ---")
    if salarioCalc, err := salario.CalcularSalario(100, 5); err != nil {
        fmt.Println("Ocurrió un error:", err)
    } else {
        fmt.Printf("El salario calculado es: %.2f€\n\n", salarioCalc)
    }
}
