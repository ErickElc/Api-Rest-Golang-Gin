package routes

import (
	"apirestgin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/html", controllers.ExibePaginaIndex)
	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/aluno/:id", controllers.AlunoPorId)
	r.GET("/aluno/cpf/:cpf", controllers.AlunoPorCPF)
	r.POST("/register-aluno", controllers.CriarAluno)
	r.PUT("/editar/:id", controllers.UpdateAluno)
	r.DELETE("/deletar/:id", controllers.DeletaAluno)
	r.NoRoute(controllers.RouteNotFound)
	r.Run(":5431")
}
