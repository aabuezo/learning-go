package main

import "testing"

func testInventory() Inventory {
	inv := NewInventory()
	inv.Add(Product{ID: 1, Name: "book", Price: 10, Stock: 10})
	inv.Add(Product{ID: 2, Name: "notebook", Price: 100, Stock: 5})
	return inv
}

func TestCartAddAndMergeItems(t *testing.T) {
	inv := testInventory()
	var cart Cart

	if err := cart.Add(CartItem{ProductID: 1, Quantity: 2}, inv); err != nil {
		t.Fatalf("Add() devolvió un error inesperado: %v", err)
	}
	if err := cart.Add(CartItem{ProductID: 1, Quantity: 3}, inv); err != nil {
		t.Fatalf("Add() devolvió un error inesperado: %v", err)
	}

	if len(cart.items) != 1 {
		t.Fatalf("el carrito tiene %d items; want 1", len(cart.items))
	}
	if cart.items[0].Quantity != 5 {
		t.Errorf("cantidad = %d; want 5", cart.items[0].Quantity)
	}
	if inv[1].Stock != 5 {
		t.Errorf("stock = %d; want 5", inv[1].Stock)
	}
}

func TestCartAddRejectsUnknownProduct(t *testing.T) {
	inv := testInventory()
	var cart Cart

	if err := cart.Add(CartItem{ProductID: 99, Quantity: 1}, inv); err == nil {
		t.Fatal("Add() debía devolver un error para un producto inexistente")
	}
	if len(cart.items) != 0 {
		t.Fatal("el producto inexistente fue agregado al carrito")
	}
}

func TestCartAddRejectsInsufficientStock(t *testing.T) {
	inv := testInventory()
	var cart Cart

	if err := cart.Add(CartItem{ProductID: 1, Quantity: 11}, inv); err == nil {
		t.Fatal("Add() debía devolver un error si no hay stock suficiente")
	}
	if len(cart.items) != 0 {
		t.Fatal("el producto fue agregado pese a no tener stock suficiente")
	}
	if inv[1].Stock != 10 {
		t.Errorf("stock = %d; want 10", inv[1].Stock)
	}
}

func TestCartChangeQuantityAdjustsStock(t *testing.T) {
	inv := testInventory()
	var cart Cart

	if err := cart.Add(CartItem{ProductID: 1, Quantity: 2}, inv); err != nil {
		t.Fatalf("Add() devolvió un error inesperado: %v", err)
	}
	if err := cart.ChangeQuantity(1, 5, inv); err != nil {
		t.Fatalf("ChangeQuantity() devolvió un error inesperado: %v", err)
	}

	if cart.items[0].Quantity != 5 {
		t.Errorf("cantidad = %d; want 5", cart.items[0].Quantity)
	}
	if inv[1].Stock != 5 {
		t.Errorf("stock = %d; want 5", inv[1].Stock)
	}
}

func TestCartChangeQuantityRestoresStockWhenQuantityDecreases(t *testing.T) {
	inv := testInventory()
	var cart Cart

	if err := cart.Add(CartItem{ProductID: 1, Quantity: 5}, inv); err != nil {
		t.Fatalf("Add() devolvió un error inesperado: %v", err)
	}
	if err := cart.ChangeQuantity(1, 2, inv); err != nil {
		t.Fatalf("ChangeQuantity() devolvió un error inesperado: %v", err)
	}

	if cart.items[0].Quantity != 2 {
		t.Errorf("cantidad = %d; want 2", cart.items[0].Quantity)
	}
	if inv[1].Stock != 8 {
		t.Errorf("stock = %d; want 8", inv[1].Stock)
	}
}

func TestCartChangeQuantityChecksOnlyAdditionalStock(t *testing.T) {
	inv := testInventory()
	var cart Cart

	if err := cart.Add(CartItem{ProductID: 1, Quantity: 2}, inv); err != nil {
		t.Fatalf("Add() devolvió un error inesperado: %v", err)
	}
	if err := cart.ChangeQuantity(1, 9, inv); err != nil {
		t.Fatalf("ChangeQuantity() devolvió un error inesperado: %v", err)
	}

	if cart.items[0].Quantity != 9 {
		t.Errorf("cantidad = %d; want 9", cart.items[0].Quantity)
	}
	if inv[1].Stock != 1 {
		t.Errorf("stock = %d; want 1", inv[1].Stock)
	}
}

func TestCartRemoveRestoresStock(t *testing.T) {
	inv := testInventory()
	var cart Cart

	if err := cart.Add(CartItem{ProductID: 1, Quantity: 3}, inv); err != nil {
		t.Fatalf("Add() devolvió un error inesperado: %v", err)
	}
	if err := cart.Remove(1, inv); err != nil {
		t.Fatalf("Remove() devolvió un error inesperado: %v", err)
	}

	if len(cart.items) != 0 {
		t.Fatal("el producto no fue eliminado del carrito")
	}
	if inv[1].Stock != 10 {
		t.Errorf("stock = %d; want 10", inv[1].Stock)
	}
}

func TestCartTotal(t *testing.T) {
	inv := testInventory()
	var cart Cart

	if err := cart.Add(CartItem{ProductID: 1, Quantity: 2}, inv); err != nil {
		t.Fatalf("Add() devolvió un error inesperado: %v", err)
	}
	if err := cart.Add(CartItem{ProductID: 2, Quantity: 1}, inv); err != nil {
		t.Fatalf("Add() devolvió un error inesperado: %v", err)
	}

	got, err := cart.Total(inv)
	if err != nil {
		t.Fatalf("Total() devolvió un error inesperado: %v", err)
	}
	if got != 120 {
		t.Errorf("Total() = %.2f; want 120", got)
	}
}

func TestCartTotalWithUnknownProduct(t *testing.T) {
	inv := testInventory()
	cart := Cart{items: []CartItem{{ProductID: 99, Quantity: 1}}}

	if _, err := cart.Total(inv); err == nil {
		t.Fatal("Total() debía devolver un error para un producto inexistente")
	}
}

func TestCartChangeQuantityUnknownProduct(t *testing.T) {
	inv := testInventory()
	var cart Cart

	if err := cart.ChangeQuantity(99, 1, inv); err == nil {
		t.Fatal("ChangeQuantity() debía devolver un error para un producto inexistente")
	}
}

func TestCartRemoveUnknownProduct(t *testing.T) {
	inv := testInventory()
	var cart Cart

	if err := cart.Remove(99, inv); err == nil {
		t.Fatal("Remove() debía devolver un error para un producto inexistente")
	}
}
