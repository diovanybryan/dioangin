// controller/controller.go
package controller

import (
    "encoding/json"
    "fmt"
    "net/http"
    "dioangin/model"
)

// StatusHandler menangani endpoint /status
func StatusHandler(w http.ResponseWriter, r *http.Request) {
    // Dapatkan status dari channel
    status := <-model.StatusChannel

    // Cetak status ke konsol
    fmt.Printf("Water: %.2f  Wind: %.2f  Status: %s\n", status.Water, status.Wind, status.Status)

    // Ubah status ke format JSON
    statusJSON, err := json.Marshal(status)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Tanggapi dengan status JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(statusJSON)
}
