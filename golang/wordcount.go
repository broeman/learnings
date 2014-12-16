package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    result := make(map[string]int)
    wordlist := strings.Fields(s)
    for _, word:= range wordlist {    	
        result[word] += 1
    }    
    return result
}

func main() {
    wc.Test(WordCount)
}