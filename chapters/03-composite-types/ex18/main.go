package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

type Inventory struct {
	Items []Product
}

// agregar producto
func (i *Inventory) Add(p Product) {
	i.Items = append(i.Items, p)
}

// buscar producto por ID
func (i *Inventory) SearchProductByID(id int) (Product, error) {
	for _, product := range i.Items {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, errors.New("product not found")
}

// actualizar stock
func (i *Inventory) Update(p Product) error {
	for idx, product := range i.Items {
		if product.ID == p.ID {
			i.Items[idx] = p
			return nil
		}
	}
	return errors.New("could not update product")
}

// listar productos
func (i Inventory) Print() {
	for _, product := range i.Items {
		fmt.Println(product)
	}
}

// calcular valor total del inventario
func (i Inventory) Total() float64 {
	total := 0.0
	for _, product := range i.Items {
		total += product.Price * float64(product.Stock)
	}
	return total
}

func NewInventory() Inventory {
	return Inventory{Items: make([]Product, 0)}
}

func main() {

	book := Product{1, "book", 36.0, 20}
	notebook := Product{2, "notebook", 999, 5}

	inventory := NewInventory()
	inventory.Add(book)
	inventory.Add(notebook)
	inventory.Print()

}
