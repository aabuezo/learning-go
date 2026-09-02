package main

import "fmt"

// crearCalculadoraPrecio devuelve una función que calcula el precio final a
// partir del impuesto y el descuento recibidos. El closure captura esos dos
// valores al crear la calculadora; el precio base se recibe y se aplica recién
// cuando se invoca la función devuelta.
func crearCalculadoraPrecio(impuesto, descuento float64) func(float64) float64 {
	return func(precio float64) float64 {
		precioFinal := (precio * (1 - descuento/100.0)) * (1 + (impuesto / 100.0))
		return precioFinal
	}
}

func main() {

	precio := 100.0
	impuesto1 := 21.0
	descuento1 := 20.0
	calc1 := crearCalculadoraPrecio(impuesto1, descuento1)
	fmt.Printf("Precio: $%.2f, Impuesto1: %.2f%%, Descuento1: %.2f%%, Precio Final: $%.2f\n",
		precio, impuesto1, descuento1, calc1(precio))

	impuesto2 := 3.0
	descuento2 := 10.0
	calc2 := crearCalculadoraPrecio(impuesto2, descuento2)
	fmt.Printf("Precio: $%.2f, Impuesto2: %.2f%%, Descuento2: %.2f%%, Precio Final: $%.2f\n",
		precio, impuesto2, descuento2, calc2(precio))
}
