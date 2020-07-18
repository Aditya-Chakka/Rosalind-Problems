package main

import(
  "fmt"
  "math/big"
  "math"
)

func main(){
  str := "UUGCGCGACGAGUUGAAAAGUGCUGAUGCCUCGGAAAUAGGGAGUAUAGGUUUCUAUAGGUGUAAGACUUUCUCCCUGAAUUCACC"
  var countA int64
  var countG int64
  var countU int64
  var countC int64
  countA = 0
  countG = 0
  countU = 0
  countC = 0
  for i := 0; i < len(str); i++{
    if string(str[i]) == "A"{
      countA++
    } else if string(str[i]) == "G"{
      countG++
    } else if string(str[i]) == "U"{
      countU++
    } else if string(str[i]) == "C"{
      countC++
    }
  }
  maxAU := int64(math.Max(float64(countA), float64(countU)))
  maxGC := int64(math.Max(float64(countG), float64(countC)))
  minAU := int64(math.Min(float64(countA), float64(countU)))
  minGC := int64(math.Min(float64(countG), float64(countC)))
  //fmt.Println(countA, countG)
  factAU := bigFactorial(maxAU, maxAU - minAU)
  factGC := bigFactorial(maxGC, maxGC - minGC)
  prod := big.NewInt(1)
  prod.Mul(factAU, factGC)
  fmt.Println(prod)
}

func factorial(i, k int64) int64 {
  var count int64 = 1
  for n := i; n > k; n-- {
    count *= n
    fmt.Println(count)
  }
  return count
}

func bigFactorial(i, k int64) *big.Int {
  count := big.NewInt(1)
  for n := i; n > k; n-- {
    count.Mul(count, big.NewInt(n))
    fmt.Println(count)
  }
  return count
}
