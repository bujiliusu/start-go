package main

import "fmt"

func lengthOfNonRepeatingSubString (s string) int {
	lastOccured := make(map[rune] int)
	start := 0
	maxlength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= 0 {
			start = lastI + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxlength
}
func main() {
	fmt.Println(lengthOfNonRepeatingSubString("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubString("你妹s都你"))
}
