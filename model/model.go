// model/model.go
package model

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Data adalah struktur untuk menyimpan data air dan angin
type Data struct {
	Water float64 `json:"water"`
	Wind  float64 `json:"wind"`
}

// Status adalah struktur untuk menyimpan status air dan angin
type Status struct {
	Water       float64 `json:"water"`
	Wind        float64 `json:"wind"`
	WaterStatus string  `json:"water_status"`
	WindStatus  string  `json:"wind_status"`
	Status      string  `json:"status"`
}

// Mu adalah mutex untuk mengamankan akses ke data
var Mu sync.Mutex

// CurrentData adalah variabel untuk menyimpan data air dan angin saat ini
var CurrentData Data

// StatusChannel adalah channel untuk mengirim status ke goroutine HTTP
var StatusChannel chan Status

// updateDataRoutine memperbarui data setiap 15 detik
func updateDataRoutine() {
	for {
		UpdateData()
		time.Sleep(15 * time.Second)
	}
}

// Init inisialisasi data dan menjalankan goroutine untuk memperbarui data secara berkala
func Init() {
	// Inisialisasi channel
	StatusChannel = make(chan Status)

	// Set nilai awal untuk data air dan angin
	CurrentData = Data{
		Water: 3.0,
		Wind:  10.0,
	}

	// Jalankan goroutine untuk memperbarui data setiap detik
	go updateDataRoutine()
}

// UpdateData memperbarui data air dan angin secara acak
func UpdateData() {
	Mu.Lock()
	defer Mu.Unlock()

	// Acak nilai air dan angin sebagai integer
	CurrentData.Water = float64(rand.Intn(100))
	CurrentData.Wind = float64(rand.Intn(20))

	// Dapatkan status dan kirim ke channel
	status := GetStatus()
	StatusChannel <- status

	fmt.Printf("Updated Data: %+v\n", CurrentData)
}

// GetStatus mendapatkan status berdasarkan kondisi air dan angin
func GetStatus() Status {
	var waterStatus, windStatus string

	// Lock mutex untuk menghindari race condition
	Mu.Lock()
	defer Mu.Unlock()

	// Tentukan status air
	if CurrentData.Water < 5 {
		waterStatus = "Safe"
	} else if CurrentData.Water >= 6 && CurrentData.Water <= 8 {
		waterStatus = "Alert"
	} else {
		waterStatus = "Danger"
	}

	// Tentukan status angin
	if CurrentData.Wind < 6 {
		windStatus = "Safe"
	} else if CurrentData.Wind >= 7 && CurrentData.Wind <= 15 {
		windStatus = "Alert"
	} else {
		windStatus = "Danger"
	}

	// Tentukan status keseluruhan
	var overallStatus string
	if waterStatus == "Danger" || windStatus == "Danger" {
		overallStatus = "Danger"
	} else if waterStatus == "Alert" || windStatus == "Alert" {
		overallStatus = "Alert"
	} else {
		overallStatus = "Safe"
	}

	return Status{
		Water:       CurrentData.Water,
		Wind:        CurrentData.Wind,
		WaterStatus: waterStatus,
		WindStatus:  windStatus,
		Status:      overallStatus,
	}
}
