package main

import "fmt"

func swap(mas_pointer *[]float64, pos1, pos2 int) {
    mas := *mas_pointer
    tmp := mas[pos1]
    mas[pos1] = mas[pos2]
    mas[pos2] = tmp
}

func partition(mas_pointer *[]float64, lo, hi int) int{
  mas := *mas_pointer
  midl := mas[(lo + hi) / 2]
  fmt.Println(midl)
  fmt.Println(mas)
  i := lo
  j := hi

  for (i <= j) {
    for (mas[i] < midl) {
      i++
    }
    for (mas[j] > midl) {
      j--
    }
    if i <= j{
          swap(&mas, i, j)
          i++
          j--
        }
          // break
          fmt.Println(mas)
  }
  midl = mas[(lo + hi) / 2]
  fmt.Println(midl)
  return j
}

// func quick_sort(mas_pointer *[]float64, low int, high int) {
//      mas := *mas_pointer
//      midl := high / 2
//      for (i:=0; i < high; i++;) {
//        if i != midl {
//           if mas[i] > mas[midl] {
//
//           }
//        }
//      }
// }

func main() {
  mas := []float64{22, 21, 419, 11, 0, 25, 22, 33, 100, 1, 0, 12}
  k := partition(&mas, 0, len(mas) - 1)
  fmt.Println(k, mas)
}
