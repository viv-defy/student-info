package handlers

import (
	"fmt"
	"strconv"
	"student-info/db"
	"student-info/models"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	students, err := db.GetAllStudents()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
	}

	c.JSON(200, students)
}

func GetStudentByID(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id parameter",
		})
		return
	}

	fmt.Printf("Id: %v\n", id)
	student, err := db.GetStudentByID(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, student)
}

func CreateStudent(c *gin.Context) {
	var students []*models.Student
	if err := c.ShouldBind(&students); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	err := db.CreateStudent(students)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func UpdateStudent(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id parameter",
		})
		return
	}

	var student models.Student
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	err = db.UpdateStudent(id, student)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, student)
}

func DeleteStudent(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id parameter",
		})
		return
	}

	err = db.DeleteStudent(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
