package main

import "testing"

func TestNewUsers(t *testing.T) {
	users := NewUsers()

	if users.users == nil {
		t.Fatal("NewUsers() debe inicializar el map")
	}
}

func TestUsersAddAndSearch(t *testing.T) {
	users := NewUsers()
	want := User{ID: 1, Name: "Pedro", Email: "pedro@mail.com", Age: 42}

	users.Add(want)

	got, err := users.Search(want.ID)
	if err != nil {
		t.Fatalf("Search() devolvió un error inesperado: %v", err)
	}

	if got != want {
		t.Errorf("Search() = %#v; want %#v", got, want)
	}
}

func TestUsersSearchNotFound(t *testing.T) {
	users := NewUsers()

	_, err := users.Search(10)
	if err == nil {
		t.Fatal("Search() debía devolver un error para un ID inexistente")
	}
}

func TestUsersUpdate(t *testing.T) {
	users := NewUsers()
	users.Add(User{ID: 1, Name: "Pedro", Age: 42})
	want := User{ID: 1, Name: "Pedro", Email: "pedro@mail.com", Age: 44}

	if err := users.Update(want); err != nil {
		t.Fatalf("Update() devolvió un error inesperado: %v", err)
	}

	got, err := users.Search(want.ID)
	if err != nil {
		t.Fatalf("Search() devolvió un error inesperado: %v", err)
	}

	if got != want {
		t.Errorf("después de Update(), Search() = %#v; want %#v", got, want)
	}
}

func TestUsersUpdateNotFound(t *testing.T) {
	users := NewUsers()

	if err := users.Update(User{ID: 10}); err == nil {
		t.Fatal("Update() debía devolver un error para un ID inexistente")
	}
}

func TestUsersDelete(t *testing.T) {
	users := NewUsers()
	users.Add(User{ID: 1, Name: "Pedro"})

	if err := users.Delete(1); err != nil {
		t.Fatalf("Delete() devolvió un error inesperado: %v", err)
	}

	if _, err := users.Search(1); err == nil {
		t.Fatal("el usuario eliminado todavía puede encontrarse")
	}
}

func TestUsersDeleteNotFound(t *testing.T) {
	users := NewUsers()

	if err := users.Delete(10); err == nil {
		t.Fatal("Delete() debía devolver un error para un ID inexistente")
	}
}
