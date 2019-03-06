package main

import "fmt"

func main() {
  target_1 := []rune("パトカー")
  target_2 := []rune("タクシー")

  for i := 0; i < len(target_1); i++ {
    fmt.Print(string(target_1[i]) + string(target_2[i]))
  }
}
