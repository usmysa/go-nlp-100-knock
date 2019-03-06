package main

import "fmt"

func main() {
  target := []rune("パタトクカシーー")

  for i := 1; i <= 7; i += 2 {
    fmt.Print(string(target[i]))
  }
}
