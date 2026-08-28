package main

type Address struct {
	Street string
	City   string
}

type Customer struct {
	ID      int
	Name    string
	Address Address
}

func main() {

	address := Address{
		Street: "123 Main St",
		City:   "Anytown",
	}

	customer := Customer{
		ID:      1,
		Name:    "John Doe",
		Address: address,
	}

	println("Customer ID:", customer.ID)
	println("Customer Name:", customer.Name)
	println("Customer Address:", customer.Address.Street, ",", customer.Address.City)
}
