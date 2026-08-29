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

	if _, err := ss.findByID(s.ID); err == nil {
		return ErrStudentExists
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

	sid, err := ss.findByID(id)
	if err != nil {
		return err
	}

	(*ss)[sid].Grades = append((*ss)[sid].Grades, g)
	return nil
}

// calcular promedio de un estudiante
func (ss Students) GetAverage(id int) (Grade, error) {
	var grades []Grade

	sid, err := ss.findByID(id)
	if err != nil {
		return 0, err
	}

	grades = ss[sid].Grades
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
func (ss Students) ListApprovedStudents() (Students, error) {
	if len(ss) == 0 {
		return nil, ErrEmptyStudents
	}

	const approved Grade = 70.0
	approvedStudents := Students{}

	for _, s := range ss {
		avg, err := ss.GetAverage(s.ID)
		if errors.Is(err, ErrNoGrades) {
			continue
		}
		if err != nil {
			return nil, err
		}

		if avg >= approved {
			approvedStudents = append(approvedStudents, s)
		}
	}

	return approvedStudents, nil
}

// buscar mejor promedio
func (ss Students) FindBestAverage() (Student, error) {
	if len(ss) == 0 {
		return Student{}, ErrEmptyStudents
	}

	var max Grade = 0.0
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

func (ss Students) findByID(id int) (int, error) {
	for i, s := range ss {
		if s.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("%w: %d", ErrStudentNotFound, id)
}

func (ss Students) Print() {
	if len(ss) == 0 {
		return
	}

	for _, s := range ss {
		fmt.Println(s)
	}
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

	var ss Students
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

	lst, err := ss.ListApprovedStudents()
	if err != nil {
		fmt.Println(err)
	}
	for _, s := range lst {
		fmt.Println(s)
	}
	fmt.Println(separator)

	s, err := ss.FindBestAverage()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Student with best average: %s\n", s.Name)
	fmt.Println(separator)
}
