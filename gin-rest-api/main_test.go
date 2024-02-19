package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-gin-api/controllers"
	"go-gin-api/database"
	"go-gin-api/models"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestVerifyGetAlunoById(t *testing.T) {
	database.Connect()
	createAlunoMock()
	defer deleteAlunoMock()
	r := setupRouter()
	r.GET("/alunos/:id", controllers.GetAlunosById)
	req, _ := http.NewRequest("GET", "/alunos/"+strconv.Itoa(ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var alunoMock models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMock)
	fmt.Println(alunoMock.Email)
	assert.Equal(t, "teste@mail.com", alunoMock.Email)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
}

func TestDeleteAluno(t *testing.T) {
	database.Connect()
	createAlunoMock()
	r := setupRouter()
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	req, _ := http.NewRequest("DELETE", "/alunos/"+strconv.Itoa(ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
}

func TestUpdateAluno(t *testing.T) {
	database.Connect()
	createAlunoMock()
	defer deleteAlunoMock()
	r := setupRouter()
	r.PATCH("/alunos/:id", controllers.UpdateAluno)
	aluno := models.Aluno{Nome: "Teste", Idade: 20, Email: "teste2@mail.com"}
	valorJson, _ := json.Marshal(aluno)
	req, _ := http.NewRequest("PATCH", "/alunos/"+strconv.Itoa(ID), bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
	var alunoMockUpdated models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMockUpdated)
	assert.Equal(t, "teste2@mail.com", alunoMockUpdated.Email)

}
