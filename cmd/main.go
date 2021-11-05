package main

import (
	"fmt"
	"net/http"

	"github.com/ValeryBMSTU/evoModeler/internal/api"
)

func main() {
	api.DevPrint()

	http.HandleFunc("/ping", api.PingHandler)
	http.HandleFunc("/", api.DoNothingHandler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
