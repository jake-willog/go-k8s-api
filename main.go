package main

import (
	"fmt"

	"net/http"

	"github.com/jake-willog/go-k8s-api.git/src/api/controller"
)

func main() {
	fmt.Println("Start.... Control Plane API")

	router := controller.NewRouter()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Err: Server")
	}
}
