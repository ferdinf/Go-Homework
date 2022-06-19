package main

import (
	"fmt"
	"go-homework/config"
)

func main() {
	fmt.Println("Starting the application...")
	config, _ := config.LoadConfiguration("config.json")
	fmt.Println(config)
}
