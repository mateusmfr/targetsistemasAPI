package main

import (
	"fmt"
	"net/http"

	"targetsistemas/internal/routes"
)

func main() {
	routes.InitializeRoutes()

	port := ":8080"
	fmt.Println("ON")

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
