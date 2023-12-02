// main.go
package main

import (
	"fmt"
	"dioangin/controller"
	"dioangin/model"
	"net/http"
)

func main() {
	// Inisialisasi model dan jalankan goroutine untuk memperbarui data setiap 15 detik
	model.Init()

	// Handler untuk endpoint /status
	http.HandleFunc("/status", controller.StatusHandler)

	// Mulai server HTTP
	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
