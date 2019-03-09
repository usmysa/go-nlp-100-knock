package main

import (
  "fmt"
)

func template(x int, y string, z float64) string{
  return fmt.Sprintf("%d", x) + "時の"+ y + "は" + fmt.Sprintf("%.1f", z)
}

func main() {
  fmt.Println(template(12, "気温", 22.4))
}
