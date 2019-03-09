package main

import (
  "fmt"
  "strings"
)

const FIRST_STR int = 0
const SECOND_STR int = 1

func isShowInitial(index int) bool {
  ary := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}

  for _, v := range ary {
    if v == index {
      return true
    }
  }

  return false
}

func main() {
  target := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
  splited := strings.Split(target, " ")

  for idx, val := range splited {
    if isShowInitial(idx + 1) == true {
      fmt.Print(string(val[FIRST_STR]))
    } else {
      fmt.Print(string(val[SECOND_STR]))
    }
  }
}
