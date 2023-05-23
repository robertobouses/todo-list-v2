package main

import (
	"fmt"

	"github.com/robertobouses/w2v-api/http"
)

func main() {
	fmt.Println("hola")
	server := http.NewServer()
	server.Run("8080")
}
