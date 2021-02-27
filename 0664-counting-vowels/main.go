package main

import (
  "fmt"
  "bufio"
  "os"
)

func scanAndMap() (string){
  var t string
  sc := bufio.NewScanner(os.Stdin)
  //var s, t string
  if sc.Scan() {
    _ = sc.Text()
  }

  if sc.Scan() {
    t = sc.Text()
  }
  return t
}

func main() {
  t := scanAndMap()
  sum := 0
  for i := 0; i < len(t); i++  {
    v := t[i:i+1]
    if (v == "a"|| v == "i" || v == "u" || v == "e" || v == "o") {
      sum++
    }
  }
  fmt.Println(sum)
  return
}
