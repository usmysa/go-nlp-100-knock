package main

import (
  "fmt"
  "strings"
  "math/rand"
	"time"
)

func sortWords(text string) string {
  splited := strings.Split(text, " ")
  if len(splited) <= 4 {
    return text
  }

  shuffledText := strings.Join(shuffle(splited[1:len(splited)-1]), " ")

  return splited[0] + " " + shuffledText + " " + splited[len(splited)-1]
}

func shuffle(words []string) []string {
  shuffled := make([]string, len(words))

  r := rand.New(rand.NewSource(time.Now().Unix()))
	for i, randIndex := range r.Perm(len(words)) {
    shuffled[i] = words[randIndex]
	}

	return shuffled
}

func main() {
  target := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

  fmt.Println(sortWords(target))
}
