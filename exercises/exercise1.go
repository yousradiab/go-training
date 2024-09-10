package exercises

import "fmt"

//importere og ekspotere function ved at starte med stort bogstav
func First() {
	fmt.Println("hej")

	//variables
	// var  bogstav rune = "b"
	// var isStupid bool = true
	// var tal int = 100

	var firstName = "Adam"

	//nemmeste måde at lave variabler
	lastName := "Warda"
	//når man redeklarer bruger man ikke :
	lastName = "Warfa"

	fmt.Println(firstName + " " + lastName)

	//Array og Slices

	numberArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numberArray)

	numberSlice := []int{1, 2, 3, 4, 5}
	//push ind i slice
	numberSlice = append(numberSlice, 6)
	fmt.Println(numberSlice)

	doubleSlice := []int{}
	//finder længde af list ved at skrive len(numberSlice)
	for i := 0; i < len(numberSlice); i++ {
		//pusher numberSlice i dobleSlice og ganger index(indholdet med 2)
		doubleSlice = append(doubleSlice, numberSlice[i]*2)
	}
	fmt.Println(doubleSlice)
}

func GetFucked(num int) (string, int) {

	if num%2 == 0 {
		return "it works", 0
	} else {
		return "go fuck", 10
	}
}
