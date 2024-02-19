package controllers

import (
	"go-gin-api/database"
	"go-gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreateAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidateAluno(&aluno); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(200, aluno)
}

func GetAlunosById(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(404, gin.H{"error": "Aluno not found"})
		return
	}
	c.JSON(200, aluno)
}

func GetAlunosByEmail(c *gin.Context) {
	var aluno models.Aluno
	email := c.Param("email")
	database.DB.Where("email = ?", email).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(404, gin.H{"error": "Aluno not found"})
		return
	}
	c.JSON(200, aluno)
}

func UpdateAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidateAluno(&aluno); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&aluno)
	c.JSON(200, aluno)
}

func DeleteAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")
	database.DB.Delete(&aluno, id)
	c.JSON(200, gin.H{"id" + id: "deleted"})
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(200, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
