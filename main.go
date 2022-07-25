package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

type User struct {
	//NOTE: attributes starting with capital letter are public and can be used outside this package

	FullName  string
	Age       int8
	BirthDate time.Time
	dni       string
}

//NOTE: This is a User class method. Commonly called as receiver
func (user *User) getFullName() string {
	return user.FullName
}

//NOTE: Interfaces are used to force that an struct implements a series of methods
type FootballPlayerInterface interface {
	Position() string
	MarketValue() float32
}

var option string

func main() {
	log.Println("Hello world!")

	var someVar string = "This is a msg"
	log.Println(someVar)
	var i int16 = 1906
	log.Println("i value is:", i)

	anotherStr := setString()
	log.Println(anotherStr)

	separator := "="
	separator = strings.Repeat(separator, 80)

	// Pointer example
	var number int = 12
	powOfTwo(&number)
	log.Println(number)

	selection := runSelector()
	switch selection {
	case "struct":
		runStructExample()
	case "map":
		runMapExample()
	case "slices":
		runSliceExample()
	case "for":
		runForLoopsExample()
	case "interface":
		runInterfacesExample()
	default:
		fmt.Println("Selected option is not valid")
	}
	/* if option == "1" {
		// Struct example
		runStructExample()
		log.Println(separator)

	} else if option == "2" {
		// Map example
		runMapExample()
		log.Println(separator)
	} else if option == "3" {
		// Slices example
		runSliceExample()
		log.Println(separator)
	} */

}

func setString() string {
	return "A string"
}

func powOfTwo(n *int) {
	*n = *n * *n
}

func runStructExample() {
	userA := User{
		FullName: "Julieta Cardenas",
		Age:      1,
	}
	log.Println(userA)
	log.Println(userA.getFullName())
}

func runMapExample() {
	// NOTE: A map is like python dicts. -> Usually maps doesn't require a pointer when passing to a func or another package
	mapA := make(map[string]string)
	mapA["brand"] = "Volkswagen"
	mapA["serie"] = "Taos"
	mapA["model"] = "2022"
	log.Println(mapA)
	log.Println(mapA["serie"])
}

func runSliceExample() {
	//NOTE: In go is normal to use slices instead of arrays
	var aSlice []string
	aSlice = append(aSlice, "Zidane")
	aSlice = append(aSlice, "Ronaldinho")
	aSlice = append(aSlice, "Ronaldo")
	log.Println(aSlice)
	log.Println(aSlice[0])
	log.Println(aSlice)
	sort.Strings(aSlice)
	log.Println(aSlice)

	numbersSlice := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbersSlice[:3])
}

func runForLoopsExample() {
	//NOTE: In go only exists for loop it can act as while or do while loop too
	animals := []string{"dog", "cat", "goat", "horse"}
	for _, animal := range animals {
		log.Println(animal)
	}
}

func runInterfacesExample() {
	zlatan := FootballPlayer{
		Name: "Zlatan Ibrahimovich",
		Age:  39,
		Club: "A.C Milan",
	}

	messi := FootballPlayer{
		Name: "Lionel Messi",
		Age:  37,
		Club: "PSG",
	}
	DisplayFootballPlayerInfo(&messi)
	DisplayFootballPlayerInfo(&zlatan)

}

type FootballPlayer struct {
	Name string
	Age  int8
	Club string
}

func (f *FootballPlayer) Position() string {
	return "Del"
}

func (f *FootballPlayer) MarketValue() float32 {
	return 100000.36
}
func DisplayFootballPlayerInfo(f FootballPlayerInterface) {
	log.Println("Position is", f.Position())
	log.Println("Market value is", f.MarketValue())
}

func runSelector() string {
	selector := make(map[string]string)
	selector["1"] = "struct"
	selector["2"] = "map"
	selector["3"] = "slice"
	selector["4"] = "for"
	selector["5"] = "interface"
	fmt.Println("1) Run struct example")
	fmt.Println("2) Run map example")
	fmt.Println("3) Run slice example")
	fmt.Println("4) Run for loops example")
	fmt.Println("5) Run interfaces example")
	fmt.Scanf("%s", &option)

	return selector[option]
}
