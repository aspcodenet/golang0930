package main

import (
	"fmt"
	"strings"
)

func main() {

	stringarna := []string{"hello", "world"}
	//for i, stringen := range stringarna {
	for _, stringen := range stringarna {
		//fmt.Println(i, stringen)
		fmt.Println(stringen)

	}

	sum := 0 // var sum = 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)

	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)

	sum = 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}

	// for - från start -> slut
	// do while - 1 och X gång
	// while - 0 och X gånger
	// while(adsdsadsa){
	//}

	//var name = "Stefan Holmberg"
	name := "Stefan Holmberg"
	// var bio = `Stefan Holmberg har programmerat sen han var 16.
	// 		Det är fortfarande det roligaste han vet.`

	if strings.Contains(name, "Stefan") {
		fmt.Println("Coolt namn")
	}

	var age uint8 = 15 // 0-255
	//var age = 15
	// var namn string = "Stefan"
	// var moms float32 = 3.12154
	// var moms2 float64 = 3.12154

	fmt.Printf("Hello World \n")
	fmt.Printf("Du är %v år gammal och heter %s ", age, namn)
}
