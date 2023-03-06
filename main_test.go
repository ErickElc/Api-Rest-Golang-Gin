package main

import (
	"apirestgin/controllers"
	"apirestgin/database"
	"apirestgin/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTests() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func TestVerifyStatusCodeAluno(t *testing.T) {
	const name string = "Erick"
	r := SetupTests()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/"+name, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")
	responseMock := `{"message":"Olá Erick, Tudo certo?"}`

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, responseMock, string(responseBody), "Deveriam ser iguais")

}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupTests()
	r.GET("/alunos", controllers.ExibeAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")

}

func TestAlunoPorId(t *testing.T) {
	database.ConectaComBancoDeDados()
	var alunoMock models.Aluno
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupTests()
	r.GET("/aluno/:id", controllers.AlunoPorId)

	req, _ := http.NewRequest("GET", "/aluno/"+strconv.Itoa(ID), nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	json.Unmarshal(response.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, response.Code)

}

func TestListandoPorCPF(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupTests()
	r.GET("/aluno/cpf/:cpf", controllers.AlunoPorCPF)

	req, _ := http.NewRequest("GET", "/aluno/cpf/12345678901", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")

}

func TestEditarAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	r := SetupTests()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r.PUT("/editar/:id", controllers.UpdateAluno)
	aluno := models.Aluno{
		Nome: "Aluno Teste", CPF: "49345678901", RG: "987654321",
	}
	valorJson, _ := json.Marshal(aluno)
	req, _ := http.NewRequest("PUT", "/editar/"+strconv.Itoa(ID), bytes.NewBuffer(valorJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var alunoMockAtualizado models.Aluno

	json.Unmarshal(response.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "49345678901", alunoMockAtualizado.CPF)
	assert.Equal(t, "987654321", alunoMockAtualizado.RG)
}

func TestDeletaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupTests()
	r.DELETE("/deletar/:id", controllers.DeletaAluno)

	req, _ := http.NewRequest("DELETE", "/deletar/"+strconv.Itoa(ID), nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Aluno Teste", CPF: "12345678901", RG: "123456789",
	}

	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno

	database.DB.Delete(&aluno, ID)
}
