package main

import (
  "fmt"
  "strings"
)

func wordNGram(target string, n int) [][]string {
  words := [][]string{}
  splited := strings.Split(target, " ")

  for i := 0; i <= len(splited) - n; i++ {
    words = append(words, splited[i:i+n])
  }

  return words
}

func charNGram(target string, n int) []string {
  words := []string{}

  for i := 0; i <= len(target) - n; i++ {
    words = append(words, target[i:i+n])
  }

  return words
}

func main() {
  target := "I am an NLPer"

  fmt.Println(wordNGram(target, 2))
  fmt.Println(charNGram(target, 2))
}
