package controllers

import (
	"apirestgin/database"
	"apirestgin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Saudacao(c *gin.Context) {
	name := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"message": "Olá " + name + ", Tudo certo?",
	})
}

func ExibeAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func AlunoPorId(c *gin.Context) {
	var alunoSelecionado models.Aluno
	idString := c.Params.ByName("id")

	database.DB.First(&alunoSelecionado, idString)
	if alunoSelecionado.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, alunoSelecionado)
}

func AlunoPorCPF(c *gin.Context) {
	var alunoSelecionado models.Aluno
	cpf := c.Param("cpf")

	database.DB.Where(&models.Aluno{CPF: cpf}).First(&alunoSelecionado)
	if alunoSelecionado.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, alunoSelecionado)
}

func CriarAluno(c *gin.Context) {
	var novoAluno models.Aluno

	if err := c.ShouldBindJSON(&novoAluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := models.ValidatorStruct(&novoAluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	database.DB.Create(&novoAluno)
	c.JSON(http.StatusOK, novoAluno)
}

func UpdateAluno(c *gin.Context) {
	var selectedAluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&selectedAluno, id)

	if err := c.ShouldBindJSON(&selectedAluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := models.ValidatorStruct(&selectedAluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	database.DB.Model(&selectedAluno).UpdateColumns(selectedAluno)

	c.JSON(http.StatusOK, selectedAluno)
}

func DeletaAluno(c *gin.Context) {
	var selectedAluno models.Aluno
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)
	database.DB.Delete(&selectedAluno, id)

	if selectedAluno.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Não foi possível deletar esse aluno!",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)

}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
