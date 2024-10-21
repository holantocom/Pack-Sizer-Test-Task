package main

import (
	"encoding/json"
	"net/http"
	"os"
	"re_partners/internal"
)

type Request struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

type Response struct {
	Result map[int]int `json:"result"`
}



// route handler function
func orderPackSizesHandler(w http.ResponseWriter, r *http.Request) {
	var payload Request
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var response Response
	response.Result = internal.CalculatePacks(payload.Target, payload.Numbers)
	json.NewEncoder(w).Encode(response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		data = []byte("Server Error")
	}
	w.Write(data)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/calculatePacks", orderPackSizesHandler)
	http.ListenAndServe(":8080", nil)
}
