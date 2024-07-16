package src

import "fmt"

type Student struct {
	StudentID   string
	StudentName string
	ListOfBooks []string
}

func (S Student) GetStudentName() string {
	return S.StudentName
}

func (S Student) SetStudentName(NewName string) {
	S.StudentName = NewName
}

func (S Student) GetListOfBooks() []string {
	return S.ListOfBooks
}

func (S Student) ListBooksBorrowedByStudent() {
	fmt.Println("Student Name: ", S.StudentName)
	fmt.Println("Student ID: ", S.StudentID)
	for index, val := range S.ListOfBooks {
		fmt.Println(index+1, ". ", val)
	}
}
