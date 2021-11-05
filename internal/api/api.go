package api

import (
	"fmt"
	"net/http"
)

func DevPrint() {
	fmt.Println("package 'api' has been attach")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s", "Что-то прилетело в PingHandler...")
	w.Write([]byte("pong"))
}

func DoNothingHandler(w http.ResponseWriter, r *http.Request) {
	return
}
