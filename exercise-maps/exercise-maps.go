/*


Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.

*/

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wdMap := make(map[string]int)

	for _, word := range strings.Fields(s) {
		if _, ok := wdMap[word]; ok {
			wdMap[word] = wdMap[word] + 1
		} else {
			wdMap[word] = 1
		}
	}
	return wdMap
}

func main() {
	wc.Test(WordCount)
}
