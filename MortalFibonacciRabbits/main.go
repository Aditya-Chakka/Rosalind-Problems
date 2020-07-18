package main

import (
  "fmt"
)

var dp []int
var n int
var m int

func main() {
//FIB RABBITS
  n = 81
  m = 16
  dp = make([]int, n)
  RabbitFib(n)
  count := 0
  for k := 0; k < m; k++ {
    //fmt.Println(dp[n-k-1])
    count += dp[n-k-1]
  }
  fmt.Println(count)
}

func RabbitFib(i int) int {
  if i == 0 {
    return 0
  }
  if i == n {
    dp[n-i] = 1
  } else if i == n - 1 {
    dp[n-i] = 0
  } else {
    for k := 2; k <= m; k++ {
      if 0 > n-i-k {
        break
      }
      //fmt.Println(dp[n-i-k], i)
      dp[n-i] += dp[n-i-k]
    }
  }
  i--
  return RabbitFib(i)
}
