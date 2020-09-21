package src

import "fmt"

type Sprter struct {
	Name string
	Sex string
	Sport string
	Grade int
}

func (a Sprter)Run()  {
	fmt.Printf("%s的运动项目是%s",a.Name,a.Sport)
}

type Car interface {
	Brand()
	Kind()
}
type Truck struct {
	Bran string
	Kin string
}

func (a Truck)Brand()  {
	fmt.Printf("%s\n",a.Bran)
}
func (a Truck)Kind()  {
	fmt.Printf("%s\n",a.Kin)
}

type ElectroCar struct {
	Bran string
	Kin string
}

func (a ElectroCar)Brand()  {
	fmt.Printf("%s\n",a.Bran)
}
func (a ElectroCar)Kind()  {
	fmt.Printf("%s\n",a.Kin)
}

type Tricycle struct {
	Bran string
	Kin  string
}

func (a Tricycle)Brand()  {
	fmt.Printf("%s\n",a.Bran)
}
func (a Tricycle)Kind()  {
	fmt.Printf("%s\n",a.Kin)
}