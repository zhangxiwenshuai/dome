package main

import "fmt"

func main() {

	var z []byte = make([]byte, 4, 4)
	z = []byte{97, 98, 99, 100}
	for i := 0; i < 4; i++ {
		fmt.Printf("%c->%d", z[i], z[i])
	}

}