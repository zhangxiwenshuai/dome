package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch1 chan int
var ch2 chan int
var z,x bool  = false,false

func main()  {
	ch1 = make(chan int,5)
	ch2 = make(chan int,5)
	e:=0
	go ran()
	go read()
	for{
		e=<-ch2
		fmt.Printf("其三次方为%d\n",e)
		if z&&x {
			break
		}
	}
	fmt.Printf("程序结束")
}
func ran()  {
	rand.Seed(time.Now().Unix())
	var w int
	for i:=0;i<100;i++  {
		w = rand.Intn(1000)
		time.Sleep(20*time.Millisecond)
		ch1 <- w
	}
	x = true
}
func read()  {
	s:=0
	for{
		s=<-ch1
		fmt.Printf("随机数为%d",s)
		ch2 <-s*s*s
		if len(ch1)==0&&x {
			break
		}
	}
	z = true
}

