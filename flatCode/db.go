package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetFur() []Furniture {
	db, err := sql.Open("postgres", "postgres://MyFlat:12345@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT color, name FROM Furniture`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	furs := make([]*Furniture, 0)
	for rows.Next() {
		fur := new(Furniture)
		err := rows.Scan(&fur.Name, &fur.Color)
		if err != nil {
			log.Fatal(err)
		}
		furs = append(furs, fur)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	convertedFurniture := make([]Furniture, len(furs))
	for i, fur := range furs {
		convertedFurniture[i] = Furniture{
			Color: fur.Color,
			Name:  fur.Name,
		}
	}
	return convertedFurniture
}

func GetApp() []Appliance {
	db, err := sql.Open("postgres", "postgres://MyFlat:12345@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT name, room FROM Appliances`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	apps := make([]*Appliance, 0)
	for rows.Next() {
		app := new(Appliance)
		err := rows.Scan(&app.Name, &app.Room)
		if err != nil {
			log.Fatal(err)
		}
		apps = append(apps, app)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	convertedAppliance := make([]Appliance, len(apps))
	for i, app := range apps {
		convertedAppliance[i] = Appliance{
			Name: app.Name,
			Room: app.Room,
		}
	}
	return convertedAppliance
}

func GetFam() []FamilyPerson {
	db, err := sql.Open("postgres", "postgres://MyFlat:12345@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT role, name, age FROM Family`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fams := make([]*FamilyPerson, 0)
	for rows.Next() {
		fam := new(FamilyPerson)
		err := rows.Scan(&fam.Role, &fam.Name, &fam.Age)
		if err != nil {
			log.Fatal(err)
		}
		fams = append(fams, fam)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	convertedFamilyPerson := make([]FamilyPerson, len(fams))
	for i, fam := range fams {
		convertedFamilyPerson[i] = FamilyPerson{
			Role: fam.Role,
			Name: fam.Name,
			Age:  fam.Age,
		}
	}
	return convertedFamilyPerson
}

func GetFla() []FlatSize {
	db, err := sql.Open("postgres", "postgres://MyFlat:12345@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT width, length, adress FROM FlatSize`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	flas := make([]*FlatSize, 0)
	for rows.Next() {
		fla := new(FlatSize)
		err := rows.Scan(&fla.Width, &fla.Length, &fla.Adress)
		if err != nil {
			log.Fatal(err)
		}
		flas = append(flas, fla)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	convertedFlatSize := make([]FlatSize, len(flas))
	for i, fla := range flas {
		convertedFlatSize[i] = FlatSize{
			Width:  fla.Width,
			Length: fla.Length,
			Adress: fla.Adress,
		}
	}
	return convertedFlatSize
}
