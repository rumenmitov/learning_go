package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
  var myArr []int 

  for _, element := range os.Args[1:] {
    intElement,_ := strconv.Atoi(element)
    myArr = append(myArr, intElement)
  }

  ch := make(chan []int)
  go mergeSort(&myArr, ch)

  fmt.Println(<-ch)
}
