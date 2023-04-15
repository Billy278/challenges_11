package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	postURL = "https://jsonplaceholder.typicode.com/posts"
)

// Struct untuk data yang akan dipost
type PostData struct {
	ValueWater int `json:"water"`
	ValueWind  int `json:"wind"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		// Generate nilai acak untuk valueWater dan valueWind antara 1-100
		valueWater := rand.Intn(100)
		valueWind := rand.Intn(100)

		// Membuat data yang akan dipost dalam format JSON
		postData := PostData{
			ValueWater: valueWater,
			ValueWind:  valueWind,
		}

		// Mengirim permintaan POST
		jsonData, err := json.Marshal(postData)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Request JSON")
		fmt.Println(string(jsonData))

		response, err := http.Post(postURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error:", err)
		}
		// Tentukan status water dan wind berdasarkan nilai acak
		statusWater := getStatusWater(valueWater)
		statusWind := getStatusWind(valueWind)

		// Menampilkan hasil response pada terminal
		fmt.Println("Response Status:", response.Status)
		fmt.Println("status water:", statusWater)
		fmt.Println("status wind :", statusWind)

		// Menunggu selama 15 detik
		time.Sleep(5 * time.Second)
	}
}

// Fungsi untuk menentukan status water berdasarkan nilai valueWater
func getStatusWater(valueWater int) string {
	if valueWater < 5 {
		return "Aman"
	} else if valueWater >= 6 && valueWater <= 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

// Fungsi untuk menentukan status wind berdasarkan nilai valueWind
func getStatusWind(valueWind int) string {
	if valueWind < 6 {
		return "Aman"
	} else if valueWind >= 7 && valueWind <= 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}
