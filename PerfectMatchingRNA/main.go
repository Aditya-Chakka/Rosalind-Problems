package main

import (
  "fmt"
  "math/big"
)



func main(){
  str := "GGCAGUUACCAAUAUGAAACUGCCAGCGACCGUAUCGAAACCCUCUUGAUCGAAUUAGGGGCCUAGCUGUGUUUCG"
  var countA int64
  var countG int64
  countA = 0
  countG = 0
  for i := 0; i < len(str); i++{
    if string(str[i]) == "A"{
      countA++
    } else if string(str[i]) == "G"{
      countG++
    }
  }
  //fmt.Println(countA, countG)
  factA := big.NewInt(factorial(countA))
  factG := big.NewInt(factorial(countG))
  prod := big.NewInt(1)
  prod.Mul(factA, factG)
  fmt.Println(prod)
}

func factorial(i int64) int64 {
  var count int64 = 1
  for n := i; n > 1; n-- {
    count *= n
  //  fmt.Println(count)
  }
  return count
}
