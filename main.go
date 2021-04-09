package main

import (
	"todos/config"
	"todos/database"
	"todos/server"
)

func main() {
	config.Load()
	database.Connect()

	server.Start()
	// fmt.Println("Name:", config.Name)
	// fmt.Println("Age:", config.Age)
	// fmt.Println("Skills:", config.Skills)
	// fmt.Println("Skills[0]:", config.Skills[0])
	// fmt.Println("Currencies:", config.Currencies)
	// fmt.Println("Currencies['bgn']:", config.Currencies["bgn"])
	// fmt.Println("Male:", config.Male)
	// fmt.Println("Birthday:", config.Birthday)
	// fmt.Println("Taxonomies:", config.Taxonomies)
}
