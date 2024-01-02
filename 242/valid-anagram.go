package main

import "reflect"

func main() {
	println(isAnagram("anagram", "nagaramp"))
}

func isAnagram(s, t string) bool {
	charS, charT := make(map[string]int), make(map[string]int)

	for _, c := range s {
		charS[string(c)] += 1
	}

	for _, c := range t {
		charT[string(c)] += 1
	}

	return reflect.DeepEqual(charS, charT)
}
