package main

import (
	"fmt"
	"go-training/exercises"
)

func main() {
	//importerer function
	// exercises.First()
	exercises.TestStudents()
	//gemmer to variabler i en linje
	getFucked, numberFuck := exercises.GetFucked(6)

	//gemmer retur værdien i en variabel.
	_, buttFuck := exercises.GetFucked(6)

	fmt.Println(numberFuck)
	fmt.Println(getFucked)
	fmt.Println(buttFuck)
}
