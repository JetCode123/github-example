package main

import "fmt" 

func getsum(until Int) Int {
  var sum Int = 0
  for i:=1 to until {
    sum += i
  }
  return sum
}

func main() {
  var n Int = 100
  fmt.Printf("1+2+..+%d=%d\n",n, getsum(n))
}
