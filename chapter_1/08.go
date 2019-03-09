package main

import (
  "fmt"
)

func cipher(plaintext string) string {
  ciphertext := ""
  start := "a"
  end := "z"

  for i := 0; i < len(plaintext); i++ {
    if start[0] <= plaintext[i] && plaintext[i] <= end[0]{
      ciphertext += string(219 - plaintext[i])
    } else {
      ciphertext += string(plaintext[i])
    }
  }

  return ciphertext
}

func main() {
  fmt.Println(cipher("Hello World"))
  fmt.Println(cipher(cipher("Hello World")))
}
