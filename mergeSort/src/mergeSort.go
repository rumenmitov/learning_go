package main

func mergeSort(arr *[]int, ch chan []int) {

  switch len(*arr) {
  case 2:
    if (*arr)[0] < (*arr)[1] {
      result := []int { (*arr)[0], (*arr)[1] }
      ch <- result
    } else {
      result := []int { (*arr)[1], (*arr)[0] }
      ch <- result
    }
  case 1:
    result := []int { (*arr)[0] }
    ch <- result
  default:
    sliceLower := (*arr)[:len(*arr)/2]
    sliceUpper := (*arr)[len(*arr)/2:]
    
    chLower := make(chan []int)
    chUpper := make(chan []int)

    go mergeSort(&sliceLower, chLower)
    go mergeSort(&sliceUpper, chUpper)
  
    sliceLowerSorted := <- chLower
    sliceUpperSorted := <- chUpper

    Outerloop:
    for _, elementLower := range sliceLowerSorted {
      for index, elementUpper := range sliceUpperSorted {
        if elementUpper >= elementLower {
          sliceUpperSorted = insertElement(&sliceUpperSorted, elementLower, index)
          continue Outerloop
        }
      } 
      sliceUpperSorted = insertElement(&sliceUpperSorted, elementLower, len(sliceUpperSorted)) 
    }
    ch <- sliceUpperSorted
  }
}

func insertElement(arr *[]int, element int, index int ) ([]int) {
  result := make([]int, len(*arr)+1)

  copy(result, (*arr)[:index])
  result[index] = element
  copy(result[index+1:], (*arr)[index:])
  return result
}
