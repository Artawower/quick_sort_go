package main

import (
  "fmt"
)

func q_sort(mas []float64) []float64{
  left_mas := []float64{}
  right_mas := []float64{}
  result_list := []float64{}
  if len(mas) > 1 {
    midl := len(mas) / 2
    for i := 0; i < len(mas); i++ {
      if mas[i] < mas[midl] && midl != i{
        left_mas = append(left_mas, mas[i])
      } else if mas[i] >= mas[midl] && midl != i {
        right_mas = append(right_mas, mas[i])
      }
    }
    left_mas = q_sort(left_mas)
    right_mas = q_sort(right_mas)
    result_list = append(left_mas, mas[midl])
    result_list = append(result_list, right_mas...)
  } else {
    result_list = mas
  }
  return result_list
}

func main() {
  var kek = []float64{77, 12, 33, 44, 50, 1, 12, 9, 4}
  l := q_sort(kek)
  fmt.Println(l)
}
