package main

import "fmt"

func main() {
	var str string = "WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA"
	fmt.Printf("%s\n", str)
	arr := make(map[string]int)
	var x string = ""
	for i := 0; i < len(str); i++ {
		x = str[i : i+1]
		if arr[x] != 0 {
			continue
		}
		for j := 0; j < len(str); j++ {
			if x == str[j:j+1] {
				arr[x]++
			}
		}
	}
	fmt.Println(arr)

 }