package main

import (
	"fmt"
	"sort"
	"strings"
)

type CartItem struct {
	ProductID int
	Quantity  int
}

type Cart struct {
	items []CartItem
}

// agregar producto al carrito
func (c *Cart) Add(ci CartItem, inv Inventory) error {
	if ci.Quantity <= 0 {
		return fmt.Errorf("la cantidad debe ser mayor que cero")
	}

	product, ok := inv[ci.ProductID]
	if !ok {
		return fmt.Errorf("el producto con ID %d no existe", ci.ProductID)
	}

	if product.Stock < ci.Quantity {
		return fmt.Errorf("no hay suficiente stock para el producto con ID %d", ci.ProductID)
	}

	if err := inv.AdjustStock(ci.ProductID, -ci.Quantity); err != nil {
		return err
	}

	for i, cartItem := range c.items {
		if cartItem.ProductID == ci.ProductID {
			c.items[i].Quantity += ci.Quantity
			return nil
		}
	}

	c.items = append(c.items, ci)
	return nil
}

// eliminar producto
func (c *Cart) Remove(id int, inv Inventory) error {
	if id <= 0 {
		return fmt.Errorf("el ID debe ser mayor que cero")
	}

	if len(c.items) == 0 {
		return fmt.Errorf("el carrito está vacío")
	}

	_, ok := inv[id]
	if !ok {
		return fmt.Errorf("el producto con ID %d no existe", id)
	}

	newCart := make([]CartItem, 0, len(c.items))
	quantityToRestore := 0
	found := false
	for _, v := range c.items {
		if v.ProductID != id {
			newCart = append(newCart, v)
		} else {
			quantityToRestore = v.Quantity
			found = true
		}
	}
	if !found {
		return fmt.Errorf("el producto con ID %d no está en el carrito", id)
	}

	if err := inv.AdjustStock(id, quantityToRestore); err != nil {
		return err
	}

	c.items = newCart
	return nil
}

// cambiar cantidad
func (c *Cart) ChangeQuantity(id, q int, inv Inventory) error {
	if q <= 0 {
		return fmt.Errorf("la cantidad debe ser mayor que cero")
	}

	product, ok := inv[id]
	if !ok {
		return fmt.Errorf("el producto con ID %d no existe", id)
	}

	for i, item := range c.items {
		if item.ProductID != id {
			continue
		}

		delta := q - item.Quantity
		if delta > product.Stock {
			return fmt.Errorf("no hay suficiente stock para el producto con ID %d", id)
		}

		if err := inv.AdjustStock(id, -delta); err != nil {
			return err
		}
		c.items[i].Quantity = q
		return nil
	}

	return fmt.Errorf("el producto con ID %d no está en el carrito", id)
}

// calcular total usando el inventario del ejercicio anterior
func (c Cart) Total(inv Inventory) (float64, error) {
	total := 0.0
	for _, cartItem := range c.items {
		product, ok := inv[cartItem.ProductID]
		if !ok {
			return 0.0, fmt.Errorf(
				"el producto con ID %d no existe", cartItem.ProductID)
		}

		total += product.Price * float64(cartItem.Quantity)
	}

	return total, nil
}

// listar carrito
func (c Cart) Print(inv Inventory) {
	fmt.Println("Carrito")
	fmt.Printf("%8s %-10s%10s%10s\n", "ID", "Name", "Price", "Quantity")

	for _, item := range c.items {
		if product, ok := inv[item.ProductID]; ok {
			fmt.Printf("%8d %-10s%10.2f%10d\n",
				product.ID, product.Name, product.Price, item.Quantity)
		} else {
			fmt.Printf("%8d %-10s%10s%10d\n",
				item.ProductID, "N/A", "N/A", item.Quantity)
		}
	}

	total, err := c.Total(inv)
	if err != nil {
		fmt.Println("Total no disponible")
		return
	}

	fmt.Printf("Total: $%21.2f\n", total)
}

// Product representa un producto en el inventario
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Inventario
type Inventory map[int]Product

// agregar producto
func (inv Inventory) Add(p Product) error {
	if p.Stock < 0 {
		return fmt.Errorf("stock inválido")
	}

	if p.Price < 0 {
		return fmt.Errorf("precio inválido")
	}

	if p.ID <= 0 {
		return fmt.Errorf("ID inválido")
	}

	if p.Name == "" {
		return fmt.Errorf("nombre inválido")
	}

	product, ok := inv[p.ID]
	if ok {
		product.Stock += p.Stock
		inv[p.ID] = product
	} else {
		inv[p.ID] = p
	}
	return nil
}

// buscar producto por ID
func (inv Inventory) Find(id int) (Product, error) {
	if product, ok := inv[id]; ok {
		return product, nil
	}
	return Product{}, fmt.Errorf("producto no encontrado")
}

// calcular valor total del inventario
func (inv Inventory) Total() float64 {
	total := 0.0
	for _, product := range inv {
		total += product.Price * float64(product.Stock)
	}

	return total
}

// actualizar stock
func (inv Inventory) AdjustStock(id, stock int) error {
	product, ok := inv[id]
	if !ok {
		return fmt.Errorf("producto no encontrado")
	}

	if product.Stock+stock >= 0 {
		product.Stock += stock
		inv[id] = product
		return nil
	}

	return fmt.Errorf(
		"stock insuficiente para el producto con ID %d", id)
}

// listar productos
func (inv Inventory) Print() {
	fmt.Println("Inventario:")
	ids := make([]int, 0, len(inv))
	for id := range inv {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for _, id := range ids {
		fmt.Println(inv[id])
	}
}

func NewInventory() Inventory {
	return make(Inventory)
}

func main() {

	book := Product{
		ID:    1000,
		Name:  "book",
		Price: 10.0,
		Stock: 20,
	}
	notebook := Product{
		ID:    1001,
		Name:  "notebook",
		Price: 1000.0,
		Stock: 5,
	}
	paper := Product{
		ID:    1002,
		Name:  "paper",
		Price: 5.0,
		Stock: 100,
	}

	separator := strings.Repeat("-", 39)

	inv := NewInventory()
	err := inv.Add(book)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = inv.Add(notebook)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = inv.Add(paper)
	if err != nil {
		fmt.Println("Error:", err)
	}
	inv.Print()
	fmt.Println(separator)

	cart := Cart{}

	item := CartItem{ProductID: book.ID, Quantity: 2}
	if err := cart.Add(item, inv); err != nil {
		fmt.Println("Error:", err)
	}
	item = CartItem{ProductID: notebook.ID, Quantity: 1}
	if err := cart.Add(item, inv); err != nil {
		fmt.Println("Error:", err)
	}
	item = CartItem{ProductID: notebook.ID, Quantity: 1}
	if err := cart.Add(item, inv); err != nil {
		fmt.Println("Error:", err)
	}
	item = CartItem{ProductID: paper.ID, Quantity: 10}
	if err := cart.Add(item, inv); err != nil {
		fmt.Println("Error:", err)
	}
	item = CartItem{ProductID: 1004, Quantity: 10}
	if err := cart.Add(item, inv); err != nil {
		fmt.Println("Error:", err)
	}

	cart.Print(inv)
	fmt.Println(separator)

	inv.Print()
	fmt.Println(separator)

	err = cart.ChangeQuantity(paper.ID, 20, inv)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = cart.Remove(notebook.ID, inv)
	if err != nil {
		fmt.Println("Error:", err)
	}
	cart.Print(inv)
	fmt.Println(separator)

	inv.Print()
	fmt.Println(separator)
}
