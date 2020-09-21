package main

import "fmt"

func main()  {
	c := []int{12,522,99,5479,123}
	fmt.Printf("%d",c)
	p :=[]*int{&c[0],&c[1],&c[2],&c[3],&c[4]}
	o := xi(p)
	fmt.Printf("%d",*o)

}
func xi(a []*int) *[]int {
	var n int = len(a)
	arr := make([]int,n,n)
	for i:=0;i<n;i++  {
		for y:=i;y<n;y++  {
			if *a[y] < *a[i] {
				*a[i],*a[y] = *a[y],*a[i]
			}
		}
		arr[i] = *a[i]
	}
	return &arr
}