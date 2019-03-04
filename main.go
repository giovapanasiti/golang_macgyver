package main

import (
	"bytes"
	"fmt"
)

// SplitSubstringN Split a string in substrings each N characters
func SplitSubstringN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

// RemoveDuplicatesInt32 remove duplicates from a int32 array
func RemoveDuplicatesInt32(a []int32) []int32 {
	r := []int32{}
	seen := map[int32]int32{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			r = append(r, val)
			seen[val] = val
		}
	}
	return r
}

func main() {
	fmt.Println("golang macgyver")
}
