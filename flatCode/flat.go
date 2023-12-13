package main

type Flat struct {
	Size       []FlatSize
	Appliances []Appliance
	Furniture  []Furniture
	Family     []FamilyPerson
}

type FlatSize struct {
	Width  float64
	Length float64
	Adress string
}

func NewFlat() Flat {
	return Flat{}
}
