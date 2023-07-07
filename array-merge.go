package main

import "fmt"

func mergeArrays(arr1 []string, arr2 []string) []string {
	merged := make([]string, 0, len(arr1)+len(arr2))
	existingNames := make(map[string]bool)

	for _, name := range arr1 {
		merged = append(merged, name)
		existingNames[name] = true
	}

	for _, name := range arr2 {
		if !existingNames[name] {
			merged = append(merged, name)
			existingNames[name] = true
		}
	}

	return merged
}

func main() {
	array1 := []string{"John", "Jane", "Mike"}
	array2 := []string{"Mike", "Alice", "Bob"}

	result := mergeArrays(array1, array2)
	fmt.Println(result)
}
