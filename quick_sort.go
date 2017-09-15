package main

import (
  "fmt"
)

func delete_slice_elem(mas []float64, index int) []float64 {
    new_slice := append(mas[0:index], mas[index + 1:len(mas)]...)
    return new_slice
}

func quick_sort(mas []float64) []float64 {
  result_list := []float64{}
  buf := []float64{}
  if len(mas) > 1 {
    midl := len(mas)/2
    result_list = append(result_list, mas[midl])

    for i:=0; i<len(mas); i++ {
      if (mas[i] < mas[midl]) && (i != midl) {
        buf = append(buf, mas[i])
        result_list = append(buf, result_list...)
        fmt.Println(result_list)
        buf = delete_slice_elem(buf, 0)
      } else if i != midl {
        result_list = append(result_list, mas[i])
      }
    }
    fmt.Println(result_list)
    left_mas := result_list[0:midl]
    right_mas := result_list[midl:len(mas)]
    left_mas = quick_sort(left_mas)
    right_mas = quick_sort(right_mas)
    result_list = append(left_mas, right_mas...)
    return result_list
  } else {
    return mas
  }

}

func main() {
  var kek = []float64{77, 12, 33, 44, 30, 1, 12, 9, 4, 555}
  kek = quick_sort(kek)
}
