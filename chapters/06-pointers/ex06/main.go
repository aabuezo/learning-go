package main

import "fmt"

type Cuenta struct {
	saldo float64
}

func (c Cuenta) Saldo() float64 {
	return c.saldo
}

func (c *Cuenta) Depositar(monto float64) {
	c.saldo += monto
}

func (c Cuenta) DepositarPorValor(monto float64) {
	c.saldo += monto
}

type Depositante interface {
	Depositar(float64)
}

func main() {
	cuenta := Cuenta{}
	cuenta.Depositar(100)

	var depositante Depositante = &cuenta
	depositante.Depositar(50)

	fmt.Println(cuenta.Saldo())
	cuenta.DepositarPorValor(50)
	fmt.Println(cuenta.Saldo())
}
