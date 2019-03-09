package main

import (
  "fmt"
)

func charNGram(target string, n int) []string {
  words := []string{}

  for i := 0; i <= len(target) - n; i++ {
    words = append(words, target[i:i+n])
  }

  return words
}

func contains(needle string, haystack []string) bool {
  for _, val := range haystack {
    if val == needle {
      return true
    }
  }

  return false
}

func unique(array []string) []string {
  unique := []string{}

  for _, val := range array {
    if ! contains(val, unique) {
      unique = append(unique, val)
    }
  }

  return unique
}

func union(a []string, b []string) []string {
  union := a

  for _, val := range b {
    union = append(union, val)
  }

  return unique(union)
}

func product(a []string, b[]string) []string {
  product := []string{}

  for _, val := range a {
    if contains(val, b) {
      product = append(product, val)
    }
  }

  return unique(product)
}

func diff(a []string, b []string) []string {
  diff := []string{}

  for _, val := range a {
    if ! contains(val, b) {
      diff = append(diff, val)
    }
  }

  return diff
}

func main() {
  target1 := "paraparaparadise"
  target2 := "paragraph"

  setX := charNGram(target1, 2)
  setY := charNGram(target2, 2)

  fmt.Println(union(setX, setY))
  fmt.Println(product(setX, setY))
  fmt.Println(diff(setX, setY))
  fmt.Println(contains("se", setX))
  fmt.Println(contains("se", setY))
}
