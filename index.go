package main

import (
	"fmt"
)

type Student struct {
    ID     int
    Name   string
    Gender string 
}

var students = make(map[int]Student) 


func addStudent(id int, name, gender string) {
    students[id] = Student{ID: id, Name: name, Gender: gender}
}


func updateStudent(id int, name, gender string) {
    if _, exists := students[id]; exists {
        students[id] = Student{ID: id, Name: name, Gender: gender}
        fmt.Println("Cập nhật sinh viên thành công.")
    } else {
        fmt.Println("Sinh viên không tồn tại.")
    }
}


func deleteStudent(id int) {
    if _, exists := students[id]; exists {
        delete(students, id)
        fmt.Println("Xóa sinh viên thành công.")
    } else {
        fmt.Println("Sinh viên không tồn tại.")
    }
}


func findStudent(id int) {
    if student, exists := students[id]; exists {
        fmt.Printf("Sinh viên tìm thấy: ID: %d, Tên: %s, Giới tính: %s\n", student.ID, student.Name, student.Gender)
    } else {
        fmt.Println("Sinh viên không tồn tại.")
    }
}


func printStudentsByGender(gender string) {
    fmt.Printf("Danh sách sinh viên giới tính %s:\n", gender)
    for _, student := range students {
        if student.Gender == gender {
            fmt.Printf("ID: %d, Tên: %s\n", student.ID, student.Name)
        }
    }
}

func main() {
    addStudent(1, "Nguyen Van A", "Nam")
    addStudent(2, "Tran Thi B", "Nữ")
    addStudent(3, "Le Van C", "Nam")

    
    fmt.Println("Danh sách sinh viên:")
    for _, student := range students {
        fmt.Printf("ID: %d, Tên: %s, Giới tính: %s\n", student.ID, student.Name, student.Gender)
    }

        findStudent(1)

    
    updateStudent(2, "Tran Thi B", "Nam")

	for _, student := range students {
		fmt.Printf("ID: %d, Tên: %s, Giới tính: %s\n",student.ID,student.Name,student.Gender);
	}
    printStudentsByGender("Nam")
    printStudentsByGender("Nữ")

       deleteStudent(3)
}