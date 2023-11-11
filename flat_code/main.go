package main

import (
	"fmt"
)

func showflatInfo(flat Flat) {
	fmt.Println("Адресс квартиры:", flat.Address)
	fmt.Println("Площадь квартиры:", flat.Size.Width*flat.Size.Length, "метров квадратных")
	fmt.Println("\n\nТехника:")
	for _, appliance := range flat.Appliances {
		fmt.Println(appliance.Name, "-", appliance.Room)
	}
	fmt.Println("\n\nМебель:")
	for _, furniture := range flat.Furniture {
		fmt.Println(furniture.Color, furniture.Name)
	}
	fmt.Println("\n\nЧлены семьи:")
	for _, person := range flat.Family {
		fmt.Println(person.Role, person.Name, ",", person.Age, "лет/года")
	}
}

func runProject() {
	flat := NewFlat("Москва, ул. Ангарская, д.57, кв.123", 6.5, 15.5)
	appliances := []Appliance{
		{"Холодильник", "Кухня"},
		{"Микроволновка", "Кухня"},
		{"Телевизор", "Спальня"},
		{"Стиральная машина", "Ванная комната"},
		{"Утюг", "Гостиная"},
	}
	furniture := []Furniture{
		{"Коричневый", "Шкаф"},
		{"Черное", "Кресло"},
		{"Белый", "Стол"},
		{"Зелёный", "Комод"},
		{"Синий", "Диван"},
	}
	familypersons := []FamilyPerson{
		{"Мама", "Надежда", 52},
		{"Папа", "Александр", 41},
		{"Дочь", "Диана", 20},
		{"Кот", "Серый", 4},
	}

	flat.Appliances = appliances
	flat.Furniture = furniture
	flat.Family = familypersons
	showflatInfo(flat)
}
