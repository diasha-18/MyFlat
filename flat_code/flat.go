package main

type Flat struct {
	Address    string
	Size       FlatSize
	Appliances []Appliance
	Furniture  []Furniture
	Family     []FamilyPerson
}

type FlatSize struct {
	Width  float64
	Length float64
}

func NewFlat(address string, width, length float64) Flat {
	size := FlatSize{Width: width, Length: length}
	return Flat{Address: address, Size: size}
}
