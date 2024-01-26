package handlers

import "github.com/gin-gonic/gin"

func GetAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func GetStudentByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func CreateStudent(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func UpdateStudent(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func DeleteStudent(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
