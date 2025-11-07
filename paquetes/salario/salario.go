package salario

import "errors"

// CalcularSalario calcula el salario mensual de un trabajador
// y valida las entradas. Devuelve (salario, error).
func CalcularSalario(horasTrabajadas, costePorHora float64) (float64, error) {
    if horasTrabajadas < 80 {
        return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 horas mensuales")
    }
    if costePorHora < 10 {
        return 0, errors.New("error: el coste por hora no puede ser inferior a 10")
    }

    salario := horasTrabajadas * costePorHora
    return salario, nil
}
