package main

import (
    "fmt"
    "pspan01/paquetes/shipping"
)

// Función que calcula el coste total de envío
func calculateTotalShippingCost(items []shipping.Shippable) float64 {
    var total float64
    for _, item := range items {
        total += item.ShippingCost()
    }
    return total
}

func main() {
    // Crear productos de ejemplo
    book := shipping.Book{Weight: 10}
    furniture := shipping.Furniture{Weight: 50, Volume: 1.85}
    digital := shipping.DigitalProduct{}

    // Crear slice con los productos
    products := []shipping.Shippable{book, furniture, digital}

    // Calcular el coste total
    total := calculateTotalShippingCost(products)

    // Mostrar resultado formateado
    fmt.Printf("El coste total de envío para todos los productos es: %.2f€\n", total)
}
