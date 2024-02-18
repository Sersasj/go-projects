package main

import (
	"go-gin-api/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
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
	if resp.Code != 200 {
		t.Errorf("Expected 200 but got %d", resp.Code)
	}
	if resp.Body.String() != "{\"message\":\"pong\"}" {
		t.Errorf("Expected {\"message\":\"pong\"} but got %s", resp.Body.String())
	}
}
