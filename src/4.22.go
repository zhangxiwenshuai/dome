package main

import (
	"fmt"
	"sync"
	"time"
)
var x sync.WaitGroup
var w sync.Mutex
var stock int = 300

func main(){
	x.Add(4)
	go zxw()
	go sale()
	go sale()
	go sale()
	x.Wait()
}
func zxw(){
	for i:=0;i<1000;i++{
		time.Sleep(8)
		w.Lock()
		if stock<300 {
			if stock==0 {
				w.Unlock()
				break
				}
			stock++
			}else{
			w.Unlock()
			break
			}
		fmt.Printf("\n 当前库存%d",stock)
		w.Unlock()
		}
	x.Done()
}
func sale(){
	for i:=0;i<1000;i++{
		time.Sleep(20)
		w.Lock()
		if stock>0 {
			stock--
			}else{
			w.Unlock()
			break
			}
		fmt.Printf("\n 当前库存%d",stock)
		w.Unlock()
		}
	x.Done()
}