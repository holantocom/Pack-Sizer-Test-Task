package main

import (
	"encoding/json"
	"net/http"
	"os"
	"sort"
)

type Request struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

type Response struct {
	Result map[int]int `json:"result"`
}

func removeDuplicatesAndNegativeAndZeroValues(arr []int) []int {
	mapVar := map[int]bool{}
	result := make([]int, 0)
	for e := range arr {
		if arr[e] < 1 {
			continue
		}
		if mapVar[arr[e]] != true {
			mapVar[arr[e]] = true
			result = append(result, arr[e])
		}
	}
	return result
}

func calculatePacks(order int, packSizes []int) map[int]int {
	// Remove all duplicates, negative and zero value
	packSizes = removeDuplicatesAndNegativeAndZeroValues(packSizes)
	if len(packSizes) == 0 {
		return map[int]int{}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	result := make(map[int]int)
	remaining := order

	// First loop. Calculate required packs amount from bigger to smaller
	for _, pack := range packSizes {
		if remaining < pack {
			continue
		}
		result[pack] = remaining / pack
		remaining = remaining % pack
	}

	// If there is anything left after packaging, add one smallest package.
	if remaining > 0 {
		result[packSizes[len(packSizes)-1]]++
	}

	// Combine smaller result into larger ones if they fit entirely
	// Iterating through pack sizes from smallest to largest
	for i := len(packSizes) - 1; i > 0; i-- {
		// Determine the current and next larger pack
		currentPack := packSizes[i]
		nextPack := packSizes[i-1]

		// Check if current pack exists in the result
		count, exist := result[currentPack]
		// Check if we have multiple smaller result that can be combined into a larger pack
		if exist && count > 1 && (count*currentPack)%nextPack == 0 {
			// Update the map with the count of the larger result
			result[nextPack] = result[nextPack] + (count*currentPack)/nextPack
			// Remove the smaller result from the map
			delete(result, currentPack)
		}
	}

	// Stack adjacent boxes if possible.
	// For example, 2 x 2000 and 1 x 1000 can be combined into 1 x 5000.
	// We go from the smaller box to the larger one to fold as much as possible.
	// Example for set [1x500, 1x1500, 1x2000, 0x4000] for value 3751: 1х500 and 1x1500 => 1x2000 and then 2x2000 => 1x4000
	for i := len(packSizes) - 1; i > 1; i-- {
		//Load current pack size and two next smaller packs
		targetSum, xPack, yPack := packSizes[i-2], packSizes[i-1], packSizes[i]
		for j := 1; j <= result[xPack]; j++ {
			for k := 0; k <= result[yPack]; k++ {
				// Trying to solve the equation x * pack(i-1) + y * pack(i) = pack(i - 2) {biggest} with calculated smaller boxes
				if targetSum == j*xPack+k*yPack {
					// If equal add new pack to the result and remove the coefficients X and Y from the smaller boxes. If they are 0, remove packs.
					result[targetSum] += 1
					if result[xPack] -= j; result[xPack] == 0 {
						delete(result, xPack)
					}
					if result[yPack] -= k; result[yPack] == 0 {
						delete(result, yPack)
					}
				}
			}
		}
	}

	// Return the map with the count of each pack
	return result
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
	response.Result = calculatePacks(payload.Target, payload.Numbers)
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
