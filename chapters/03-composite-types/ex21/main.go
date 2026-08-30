package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
)

var (
	ErrContactAlreadyExists = errors.New("contact already exists in contact list")
	ErrContactNotFound      = errors.New("the provided email does not exist")
	ErrEmptyContactList     = errors.New("the Contact List is empty")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidSort          = errors.New("invalid sorting code")
)

type SortBy int

const (
	SortByName SortBy = iota
	SortByEmail
	SortByPhone
)

var emailPattern = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

type Contact struct {
	Name  string
	Email string
	Phone string
}

type contactList map[string]Contact

func validateEmail(email string) error {
	if !emailPattern.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}

// agregar contacto;
func (cl contactList) Add(c Contact) error {
	if err := validateEmail(c.Email); err != nil {
		return err
	}

	_, ok := cl[c.Email]
	if ok {
		return ErrContactAlreadyExists
	}

	cl[c.Email] = c

	return nil
}

// buscar por email;
func (cl contactList) FindByEmail(email string) (Contact, error) {
	if err := validateEmail(email); err != nil {
		return Contact{}, err
	}

	contact, ok := cl[email]
	if !ok {
		return Contact{}, ErrContactNotFound
	}

	return contact, nil
}

// actualizar teléfono;
func (cl contactList) UpdatePhone(email, phone string) error {
	if err := validateEmail(email); err != nil {
		return err
	}

	contact, ok := cl[email]
	if !ok {
		return ErrContactNotFound
	}

	contact.Phone = phone
	cl[email] = contact

	return nil
}

// borrar contacto;
func (cl contactList) Delete(email string) error {
	if err := validateEmail(email); err != nil {
		return err
	}

	if _, ok := cl[email]; !ok {
		return ErrContactNotFound
	}

	delete(cl, email)
	return nil
}

// listar todos.
func (cl contactList) List(sortBy SortBy) ([]Contact, error) {
	if len(cl) == 0 {
		return []Contact{}, ErrEmptyContactList
	}

	contacts := make([]Contact, 0, len(cl))
	for _, contact := range cl {
		contacts = append(contacts, contact)
	}

	if err := sortContacts(contacts, sortBy); err != nil {
		return []Contact{}, err
	}

	return contacts, nil
}

// ordenar
func sortContacts(contacts []Contact, sortBy SortBy) error {
	switch sortBy {
	case SortByName:
		sort.Slice(contacts, func(i, j int) bool {
			return contacts[i].Name < contacts[j].Name
		})
	case SortByEmail:
		sort.Slice(contacts, func(i, j int) bool {
			return contacts[i].Email < contacts[j].Email
		})
	case SortByPhone:
		sort.Slice(contacts, func(i, j int) bool {
			return contacts[i].Phone < contacts[j].Phone
		})
	default:
		return ErrInvalidSort
	}

	return nil
}

func NewContactList() contactList {
	return make(contactList)
}

func main() {
	john := Contact{
		Name:  "john",
		Email: "john@mail.com",
		Phone: "1 555 12345",
	}
	jane := Contact{
		Name:  "jane",
		Email: "jane@mail.com",
		Phone: "1 555 12346",
	}
	patrick := Contact{
		Name:  "patrick",
		Email: "patrick@mail.com",
		Phone: "1 555 12347",
	}
	mark := Contact{
		Name:  "mark",
		Email: "mark@mail.com",
		Phone: "N/A",
	}

	cl := NewContactList()

	if err := cl.Add(john); err != nil {
		fmt.Println(err)
	}
	if err := cl.Add(jane); err != nil {
		fmt.Println(err)
	}
	if err := cl.Add(patrick); err != nil {
		fmt.Println(err)
	}
	if err := cl.Add(mark); err != nil {
		fmt.Println(err)
	}

	contacts, err := cl.List(SortByEmail)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, c := range contacts {
			fmt.Println(c)
		}
	}

	c, err := cl.FindByEmail("jane@mail.com")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}

	if err := cl.UpdatePhone("jane@mail.com", "1 555 12344"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("updated phone successfully", cl["jane@mail.com"])
	}

	if err := cl.Delete("jane@mail.com"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("entry deleted successfully")
	}

	contacts, err = cl.List(SortByName)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, c := range contacts {
			fmt.Println(c)
		}
	}
}
