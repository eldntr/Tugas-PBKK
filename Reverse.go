package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func reverse() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a string: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    reversed := reverseWords(input)
    fmt.Println("Reversed words:", reversed)
}

func reverseWords(s string) string {
    words := strings.Fields(s)
    for i, word := range words {
        words[i] = reverseString(word)
    }
    return strings.Join(words, " ")
}

func reverseString(s string) string {
    runes := []rune(s)
    n := len(runes)
    for i := 0; i < n/2; i++ {
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }
    return string(runes)
}

func main() {
    reverse()
}