package main

import (
	"fmt"
	"github.com/mini-douyin/middleware/auth"
	"net/http"
)

func main() {
	http.HandleFunc("/auth", auth.Auth)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
