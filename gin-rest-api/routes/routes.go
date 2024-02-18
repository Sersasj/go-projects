package routes

import (
	"go-gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.GET("/alunos", controllers.GetAlunos)
	r.POST("/alunos", controllers.CreateAluno)
	r.GET("/alunos/:id", controllers.GetAlunosById)
	r.GET("/alunos/email/:email", controllers.GetAlunosByEmail)
	r.PATCH("/alunos/:id", controllers.UpdateAluno)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.Run()
}
