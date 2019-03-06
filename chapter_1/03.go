package main

import (
  "fmt"
  "strings"
)

func main() {
  target := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
  splited := strings.Split(target, " ")

  for i := 0; i < len(splited); i++ {
    fmt.Print(string(splited[i][0]))
  }
}
