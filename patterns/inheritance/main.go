package main

import "fmt"

type TransporterCommon struct {
	NumWheels int
	Mileage   int
	TankVol   int
}

// Transporter is a type interface. Note that in Go, interface names typically ends with 'er'.
type Transporter interface {
	Travel() int
}

// TransporterConstraint is a type constraint interface using as a
// (Generic) constraint parameter.
type TransporterConstraint interface {
	Car | Motorcycle
}

type Car struct {
	TransporterCommon
}

type Motorcycle struct {
	TransporterCommon
}

type Bicycle struct {
	TransporterCommon
}

func (tc TransporterCommon) Travel() int {
	return tc.Mileage * tc.TankVol
}

// Travel in "subclass" Bicycle doesn't override the base class Travel method.

func (b Bicycle) Travel() int {
	return b.Mileage
}

//func TravelInTerrain[T TransporterConstraint](terrain string) int {
//
//}

func main() {
	vehicles := []Transporter{
		Car{TransporterCommon{
			Mileage:   25,
			TankVol:   12,
			NumWheels: 4,
		}},
		Motorcycle{TransporterCommon{
			Mileage:   40,
			TankVol:   3,
			NumWheels: 2,
		}},
		Motorcycle{TransporterCommon{
			Mileage:   100,
			NumWheels: 2,
		}},
	}

	for _, v := range vehicles {
		fmt.Println(v.Travel())
	}
}
