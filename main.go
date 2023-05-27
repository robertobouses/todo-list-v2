package main

import (
	"fmt"

	"github.com/robertobouses/todo-list-v2/http"
)

func main() {
	fmt.Println("hola")
	server := http.NewServer()
	server.Run("8081")
}
