package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  target := stdin.Text()

  one_count := 0
  for _, v := range strings.Split(target, " ") {
    if v == "1" {
      one_count++
    }
    if one_count == 2 {
      fmt.Println(1)
      return
    }
  }
  fmt.Println(2)
  return
}
