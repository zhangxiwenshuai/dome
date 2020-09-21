package main

import "fmt"

func main() {
	z1:=Truck{name:"解放者一号"}
	x2:=electric{name:"特斯拉电动车"}
	w3:=Tricycle{name:"时风三轮车"}
	z1.drive()
	x2.drive()
	w3.drive()
}

type Vehicle interface{
	brand()    string
	drive()    string
}
type Truck struct {
	name string
}

func (z Truck)brand() string  {
	return z.name
}
func (z Truck) drive() {
	fmt.Println("解放者一号")
}
type electric struct {
	name string
}
func (x electric)brand()string  {
	return x.name
}
func (x electric) drive() {
	fmt.Println("特斯拉电动车")
}
type Tricycle struct {
	name string
}
func (w Tricycle)brand()string  {
	return w.name
}
func (w Tricycle) drive() {
	fmt.Println("时风三轮车")
}