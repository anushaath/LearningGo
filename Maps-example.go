#Exercise: Maps
#Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

#You might find strings.Fields helpful.


package main

import (
	"golang.org/x/tour/wc"
	"strings"
	
)

func WordCount(s string) map[string]int {
	words:= strings.Fields(s)
	mapping:= make(map[string]int)
	for _,v := range words {
		//fmt.Println(i)
		if _, ok:= mapping[v]; ok{
			mapping[v] += 1
		}else{
			mapping[v]=1
		}
	}
	return mapping
}

func main() {
	wc.Test(WordCount)
}
