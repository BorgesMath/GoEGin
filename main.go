package main

import (
	"github.com/BorgesMath/GoEGin/database"
	"github.com/BorgesMath/GoEGin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()

}
