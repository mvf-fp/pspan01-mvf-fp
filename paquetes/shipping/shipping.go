package shipping

// Interfaz pública Shippable
type Shippable interface {
    ShippingCost() float64
}

// --- Structs públicos que implementan la interfaz ---

type Book struct {
    Weight float64 // en kg
}

type Furniture struct {
    Weight float64 // en kg
    Volume float64 // en m³
}

type DigitalProduct struct{}

// --- Implementaciones del método ShippingCost() ---

func (b Book) ShippingCost() float64 {
    return b.Weight * 5
}

func (f Furniture) ShippingCost() float64 {
    return f.Weight*3 + (f.Volume / 0.01)
}

func (d DigitalProduct) ShippingCost() float64 {
    return 0
}
