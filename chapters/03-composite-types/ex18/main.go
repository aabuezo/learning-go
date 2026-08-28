package main

import (
	"errors"
	"fmt"
	"sort"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

type Inventory map[int]Product

// agregar producto
func (inv Inventory) Add(product Product) {
	inv[product.ID] = product
}

// buscar producto por ID
func (inv Inventory) SearchProductByID(id int) (Product, error) {
	if product, ok := inv[id]; ok {
		return product, nil
	}
	return Product{}, errors.New("product not found")
}

// actualizar stock
func (inv Inventory) Update(id, stock int) error {
	if product, ok := inv[id]; ok {
		product.Stock = stock
		inv[id] = product
		return nil
	}
	return errors.New("product not found")
}

// listar productos
func (inv Inventory) Print() {
	ids := make([]int, 0, len(inv))
	for id := range inv {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for _, id := range ids {
		fmt.Println(inv[id])
	}
}

// calcular valor total del inventario
func (inv Inventory) Total() float64 {
	total := 0.0
	for _, product := range inv {
		total += product.Price * float64(product.Stock)
	}
	return total
}

func NewInventory() Inventory {
	return make(Inventory)
}

func main() {

	book := Product{
		ID:    1,
		Name:  "book",
		Price: 36.0,
		Stock: 20,
	}
	notebook := Product{
		ID:    2,
		Name:  "notebook",
		Price: 999,
		Stock: 5,
	}
	paper := Product{
		ID:    3,
		Name:  "paper",
		Price: 5.5,
		Stock: 100,
	}

	inv := NewInventory()
	inv.Add(book)
	inv.Add(notebook)
	inv.Add(paper)
	inv.Print()
	fmt.Println("----------------------")

	// buscar id que no existe
	_, err := inv.SearchProductByID(10)
	if err != nil {
		fmt.Println(err)
	}

	// buscar un id que si existe
	prod, err := inv.SearchProductByID(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(prod)
	fmt.Println("----------------------")

	err = inv.Update(2, 10)
	if err != nil {
		fmt.Println(err)
	}
	inv.Print()
	fmt.Println("----------------------")

	fmt.Println("Total Inventory: $", inv.Total())
}
