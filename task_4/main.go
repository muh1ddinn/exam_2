package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	shortest := strs[0]
	for _, str := range strs {
		if len(str) < len(shortest) {
			shortest = str
		}
	}

	for i, char := range shortest {
		for _, str := range strs {
			if str[i] != byte(char) {
				return shortest[:i]
			}
		}
	}

	return shortest
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(strs)
	fmt.Println("Longest common prefix:", prefix) 
}
