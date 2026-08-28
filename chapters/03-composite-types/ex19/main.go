package main

type CartItem struct {
	ProductID int
	Quantity  int
}

type Cart struct {
	Items []CartItem
}

// agregar producto al carrito
func (c *Cart) Add(ci CartItem) {
	c.Items = append(c.Items, ci)
}

// eliminar producto
func (c *Cart) Remove(id int) {
	newCart := make([]CartItem, 0)
	for _, v := range c.Items {
		if v.ProductID != id {
			newCart = append(newCart, v)
		}
	}
	c.Items = newCart
}

// cambiar cantidad
func (c *Cart) ChangeQuantity(id, q int) {
	for i, v := range c.Items {
		if v.ProductID == id {
			c.Items[i].Quantity = q
		}
	}
}

// calcular total usando el inventario del ejercicio anterior

func main() {
}
