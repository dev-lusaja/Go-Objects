package main

import "fmt"

type Animal struct {
	Tipo   string
	Idioma string
}

//Funcion asociada a la estructura Animal
func (a Animal) Comunicarse(val string) {
	Print(val)
}

//Structura Perro que hereda atributos de la estructura Animal
type Perro struct {
	Animal
	Nombre string
	Raza   string
	Edad   int
}

//Funcion asociada a la estructura Perro
func (p Perro) Ladrar() string {
	return "Guauu!!"
}

func Print(msg interface{}) {
	fmt.Println(msg)
}

func main() {
	mascota := new(Perro)
	(*mascota).Animal.Idioma = "Ladra"
	(*mascota).Animal.Tipo = "Perro"
	(*mascota).Nombre = "Boby"
	(*mascota).Edad = 7
	(*mascota).Raza = "Boxer"
	Print(*mascota)
	(*mascota).Animal.Comunicarse((*mascota).Ladrar())
}
