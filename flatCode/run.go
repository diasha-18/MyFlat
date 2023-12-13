package main

import (
	"fmt"
)

func ShowFlatInfo(f Flat) {
	fmt.Printf("Адрес квартиры: \n")
	for _, fla := range f.Size {
		fmt.Printf(" %s \n", fla.Adress)
	}

	fmt.Println("\nПлощадь квартиры:")
	for _, fla := range f.Size {
		fmt.Printf(" %f квадратных метров\n", fla.Length*fla.Width)
	}

	fmt.Println("\nСемья:")
	for _, fam := range f.Family {
		fmt.Printf("%s %s. Возраст - %d лет/год(а).\n", fam.Role, fam.Name, fam.Age)
	}

	fmt.Println("\nМебель:")
	for _, fur := range f.Furniture {
		fmt.Printf("%s %s\n", fur.Name, fur.Color)
	}

	fmt.Println("\nТехника:")
	for _, app := range f.Appliances {
		fmt.Printf("%s. Местоположение - %s\n", app.Name, app.Room)
	}
}

func runProject() {
	flat := NewFlat()
	furs := GetFur()
	apps := GetApp()
	fams := GetFam()
	flas := GetFla()

	flat.Furniture = furs
	flat.Appliances = apps
	flat.Family = fams
	flat.Size = flas

	ShowFlatInfo(flat)
}
