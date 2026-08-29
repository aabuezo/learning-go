package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrStudentNotFound = errors.New("student not found")
	ErrStudentExists   = errors.New("student already exists")
	ErrNoGrades        = errors.New("student has no grades")
	ErrEmptyStudents   = errors.New("there are no students")
)

type Student struct {
	ID     int
	Name   string
	Grades []Grade
}

type Students []Student

type Grade float64

// agregar estudiante
func (ss *Students) AddStudent(s Student) error {
	if ss == nil {
		return fmt.Errorf("you need to create Students first")
	}

	for _, student := range *ss {
		if s.ID == student.ID {
			return ErrStudentExists
		}
	}

	*ss = append(*ss, s)
	return nil
}

// agregar nota
func (ss *Students) AddGrade(id int, g Grade) error {
	if ss == nil {
		return fmt.Errorf("you need to create Students first")
	}

	if len(*ss) == 0 {
		return ErrEmptyStudents
	}

	if g < 0 || g > 100 {
		return fmt.Errorf("grade must be between 0 and 100")
	}

	for i, student := range *ss {
		if student.ID == id {
			(*ss)[i].Grades = append((*ss)[i].Grades, g)
			return nil
		}
	}

	return fmt.Errorf("%w: %d", ErrStudentNotFound, id)
}

// calcular promedio de un estudiante
func (ss Students) GetAverage(id int) (Grade, error) {
	var grades []Grade
	present := false
	for _, s := range ss {
		if s.ID == id {
			grades = s.Grades
			present = true
		}
	}

	if !present {
		return 0.0, fmt.Errorf("%w: %d", ErrStudentNotFound, id)
	}

	if len(grades) == 0 {
		return 0.0, fmt.Errorf("%w: %d", ErrNoGrades, id)
	}

	total := Grade(0.0)
	for _, v := range grades {
		total += v
	}
	return total / Grade(len(grades)), nil
}

// listar estudiantes aprobados
func (ss Students) ListApprovedStudents() error {
	if len(ss) == 0 {
		return ErrEmptyStudents
	}

	const approved float64 = 70.0

	for _, s := range ss {
		avg, err := ss.GetAverage(s.ID)
		if err != nil {
			if errors.Is(err, ErrNoGrades) {
				continue
			}
			return err
		}

		if avg >= Grade(approved) {
			fmt.Printf("%s approved with average: %.2f\n", s.Name, avg)
		}
	}

	return nil
}

// buscar mejor promedio
func (ss Students) FindBestAverage() (Student, error) {
	if len(ss) == 0 {
		return Student{}, ErrEmptyStudents
	}

	max := Grade(0.0)
	idx := 0
	found := false
	for i, s := range ss {
		avg, err := ss.GetAverage(s.ID)
		if errors.Is(err, ErrNoGrades) {
			continue
		}
		if err != nil {
			return Student{}, err
		}

		if avg >= max {
			max = avg
			idx = i
		}
		found = true
	}

	if !found {
		return Student{}, fmt.Errorf("there are no students with grades")
	}

	return ss[idx], nil
}

func (ss Students) Print() error {
	if len(ss) == 0 {
		return ErrEmptyStudents
	}

	for _, s := range ss {
		fmt.Println(s)
	}
	return nil
}

func NewStudents() Students {
	return make(Students, 0)
}

func main() {

	john := Student{
		ID:     1,
		Name:   "John",
		Grades: []Grade{},
	}
	jane := Student{
		ID:     2,
		Name:   "Jane",
		Grades: []Grade{},
	}
	patrick := Student{
		ID:     3,
		Name:   "Patrick",
		Grades: []Grade{},
	}

	ss := NewStudents()
	if err := ss.AddStudent(jane); err != nil {
		fmt.Println(err)
	}
	if err := ss.AddStudent(patrick); err != nil {
		fmt.Println(err)
	}
	if err := ss.AddStudent(john); err != nil {
		fmt.Println(err)
	}
	ss.Print()

	separator := strings.Repeat("-", 35)
	fmt.Println(separator)

	grades := []struct {
		studentID int
		grades    []Grade
	}{
		{john.ID, []Grade{80, 90, 100}},
		{jane.ID, []Grade{75, 75}},
		{patrick.ID, []Grade{40, 50, 80}},
	}

	for _, group := range grades {
		for _, grade := range group.grades {
			if err := ss.AddGrade(group.studentID, grade); err != nil {
				fmt.Println(err)
			}
		}
	}
	ss.Print()
	fmt.Println(separator)

	avg, err := ss.GetAverage(john.ID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s has average: %.2f\n", john.Name, avg)
	}

	avg, err = ss.GetAverage(jane.ID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s has average: %.2f\n", jane.Name, avg)
	}

	avg, err = ss.GetAverage(patrick.ID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s has average: %.2f\n", patrick.Name, avg)
	}

	fmt.Println(separator)

	err = ss.ListApprovedStudents()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(separator)

	s, err := ss.FindBestAverage()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Student with best average: %s\n", s.Name)
	fmt.Println(separator)
}
