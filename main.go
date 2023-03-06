package main

import (
	"apirestgin/database"
	"apirestgin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
