package main

import (
	"go-gin-api/controllers"
	"go-gin-api/database"
	"go-gin-api/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r.GET("/ping", ping)
	// r.GET("/alunos", getAlunos)
	// r.POS T("/alunos", createAluno)
	// r.GET("/alunos/:id", getAlunosById)
	// r.GET("/alunos/email/:email", getAlunosByEmail)
	// r.PATCH("/alunos/:id", updateAluno)
	// r.DELETE("/alunos/:id", deleteAluno)
	return r
}

func TestVerifyPing(t *testing.T) {
	r := setupRouter()
	r.GET("/ping", controllers.Ping)
	req, _ := http.NewRequest("GET", "/ping", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
	respMock := `{"message":"pong"}`
	respBody, _ := io.ReadAll(resp.Body)
	assert.Equal(t, respMock, string(respBody), "Expected pong")
}

func createAlunoMock() {
	aluno := models.Aluno{Nome: "Teste", Idade: 20, Email: "teste@mail.com"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func deleteAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}
func TestVerifyGetAlunos(t *testing.T) {
	database.Connect()
	createAlunoMock()
	defer deleteAlunoMock()
	r := setupRouter()
	r.GET("/alunos", controllers.GetAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")

}

func TestVerifyGetAlunoByEmail(t *testing.T) {
	database.Connect()
	createAlunoMock()
	defer deleteAlunoMock()
	r := setupRouter()
	r.GET("/alunos/email/:email", controllers.GetAlunosByEmail)
	req, _ := http.NewRequest("GET", "/alunos/email/teste@mail.com", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
}
