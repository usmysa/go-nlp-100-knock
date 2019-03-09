package main

import "fmt"

func main() {
  target := "stressed"

  for i := len(target) - 1; i > 0; i-- {
    fmt.Print(string(target[i]))
  }
}
