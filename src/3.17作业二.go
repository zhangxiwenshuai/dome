package main

import "fmt"

func main() {
    var z [20]int = [20]int{3, 9, 10, 11, 8, 4, 9, 1, 20, 10, 11, 21, 19, 3, 8, 4, 1, 10, 20, 21}
    fmt.Printf("%d\n", z)
    arr2 := make(map[int]int)
    var w int = 0
    for i := 0; i < len(z); i++ {
     w = z[i]
    if arr2[w] != 0 {
   continue
   }
    for j := 0; j < len(z); j++ {
   if w == z[j] {
    arr2[w]++
        }
     }
   }
   fmt.Println(arr2)
}