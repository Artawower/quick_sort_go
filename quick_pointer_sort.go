package main

import (
  "fmt"
)


func partition(mas_pointer *[]float64, l int, h int) int {
  mas := *mas_pointer
  midl := mas[(l + h) / 2]
  for l <= h {
    for mas[l] < midl {
      l++
    }
    for mas[h] > midl {
      h--
    }
    if l <= h {
      tmp := mas[l]
      mas[l] = mas[h]
      mas[h] = tmp
      l++
      h--
    }
  }
  fmt.Printf("l - %d, h - %d\n", l, h)
  fmt.Println(mas)
  return l
}

func quick_sort(mas_pointer *[]float64, l int, h int) {
  if l < h{
    p := partition(mas_pointer, l, h)
    quick_sort(mas_pointer, l, p - 1)
    quick_sort(mas_pointer, p, h)
  }
}

func main() {
  mas := []float64{12, 11, 1923, 8942, 123, 941, -643}
  fmt.Println("Массив до начала сортировки\n", mas)
  quick_sort(&mas, 0, len(mas) - 1)
  // partition(&mas, 0, len(mas) - 1)
  fmt.Println("Массив после начала сортировки\n", mas)
  // partition(&mas, 0, len(mas) - 1)
  // fmt.Println(p)
}
