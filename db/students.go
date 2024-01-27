package db

import (
	"student-info/config"
	"student-info/models"
	"time"
)

func GetAllStudents() ([]*models.Student, error) {
	var students []*models.Student
	result := config.Db.Find(&students)
	return students, result.Error
}

func GetStudentByID(id int) (models.Student, error) {
	var student models.Student
	result := config.Db.Find(&student, id)
	return student, result.Error
}

func CreateStudent(student models.Student) error {
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	result := config.Db.Create(&student)
	return result.Error
}

func UpdateStudent(id int, student models.Student) error {
	result := config.Db.Model(&models.Student{}).Where("id = ?", id).Updates(models.Student{Name: student.Name, Email: student.Email, UpdatedAt: time.Now()})
	return result.Error
}

func DeleteStudent(id int) error {
	result := config.Db.Delete(&models.Student{}, id)
	return result.Error
}
