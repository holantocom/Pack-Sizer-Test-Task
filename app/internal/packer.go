package internal

import (
	"math"
	"sort"
)

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

// greatest common divisor (GCD) via Euclidean algorithm
func getGCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Creates map of possible packs that equal to order amount
func findPackMatrix(possiblePacks [][]int, currentPacks []int, index int, packSizes []int, order int) [][]int {
	if order == 0 {
		return append(possiblePacks, currentPacks)
	}
	for i := index; i < len(packSizes); i++ {
		if order-packSizes[i] < 0 {
			continue
		}
		possiblePacks = findPackMatrix(possiblePacks, append(currentPacks, packSizes[i]), i, packSizes, order-packSizes[i])
	}
	return possiblePacks
}

func findArraySum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func CalculatePacks(order int, packSizes []int) map[int]int {
	// Remove all duplicates, negative and zero value
	packSizes = removeDuplicatesAndNegativeAndZeroValues(packSizes)
	if len(packSizes) == 0 {
		return map[int]int{}
	}

	sort.Ints(packSizes)
	// If order smaller than minimum pack - return only 1 pack
	if order < packSizes[0] {
		return map[int]int{packSizes[0]: 1}
	}

	// If only one element in the list - return it
	if len(packSizes) == 1 {
		return map[int]int{packSizes[0]: int(math.Ceil(float64(order) / float64(packSizes[0])))}
	}

	// Calculating Greatest Common Divisor for requested pack.
	commonDivider := getGCD(packSizes[0], packSizes[1])
	for i := 2; i < len(packSizes); i++ {
		commonDivider = getGCD(commonDivider, packSizes[i])
	}

	// Round order to the nearest number that divisible by GCD
	order += (commonDivider - order%commonDivider) % commonDivider

	// Calculate all possible packs than can be created
	var possiblePacks [][]int
	var currentPacks []int
	possiblePacks = findPackMatrix(possiblePacks, currentPacks, 0, packSizes, order)
	// If no packs exists for requested order(typically when GCD=1), trying to add +1 to find packs
	for possiblePacks == nil {
		order += 1
		possiblePacks = findPackMatrix(possiblePacks, currentPacks, 0, packSizes, order)
	}

	// Loop through possible packs to find the smallest pack
	smallestLen := len(possiblePacks[0])
	smallestPacks := possiblePacks[0]
	for _, packs := range possiblePacks {
		// Calculating the sum for a situation where the values differ by more than 10 times
		// and the smallest array is not always the best 751 <= [5, 12, 250]
		if len(packs) < smallestLen && findArraySum(packs) == order {
			smallestLen = len(packs)
			smallestPacks = packs
		}
	}

	// Convert it to the map
	result := make(map[int]int)
	for _, pack := range smallestPacks {
		result[pack] += 1
	}
	return result
}
