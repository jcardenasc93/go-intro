package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/jcardenasc93/go-intro/helpers"
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

//NOTE: Channels are usefull to pass information between different parts of the program
//  without access vars values. This is extremely useful when sharing info between packages
func CalcRandValue(intChan chan int) {
	const n = 10
	randValue := helpers.GenerateRandInt(n)
	log.Println(randValue)
	//NOTE: chan <- is the way to pass value to a channel
	intChan <- randValue

}

var option string

type Animal struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	IsAdopted bool   `json:"is_adopted"`
}

func checkErrorNilValue(err *error, msg string) {
	if *err != nil {
		log.Panic(msg, "\n", *err)
	}

}

func jsonHandler() {

	aJson := `
  [
      {
        "name": "Tokyo",
        "type": "cat",
        "is_adopted": true
      },
      {
        "name": "Bruno",
        "type": "dog",
        "is_adopted": false
      }
  ]
  `
	var unmarshalledJson []Animal
	//NOTE: Unmarshal is the process to convert a plain json bytes to struct
	err := json.Unmarshal([]byte(aJson), &unmarshalledJson)
	checkErrorNilValue(&err, "Cannot unmarhall json bytes. Maybe malformed json")
	log.Printf("Unmarhaled json: %v\n", unmarshalledJson)

	//NOTE: Marshal is the way to go in order to convert struct to plain json bytes
	animals, err := json.MarshalIndent(unmarshalledJson, "", "    ")
	checkErrorNilValue(&err, "Cannot marhall given struct")
	log.Println("Marshalled struct:\n", string(animals))
}

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
	case "packages":
		runPackagesExample()
	case "channels":
		runChannelsExample()
	case "jsonHandler":
		jsonHandler()
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

func runPackagesExample() {
	testCar := helpers.Car{
		Brand: "BMW",
		Model: 2014,
		Color: "#000",
	}
	log.Println("The following Car type variable comes from helper package")
	log.Println(testCar)
}

func runChannelsExample() {
	intChan := make(chan int)
	//NOTE: Run the func in a gorutine
	go CalcRandValue(intChan)
	//NOTE: <- channel is the way to read channel value
	channelValue := <-intChan
	log.Println("Current channel value is", channelValue)

	//NOTE: defer helps to run this always before living the func. It doesn't matter where is putted
	defer close(intChan)

}
func runSelector() string {
	selector := make(map[string]string)
	selector["1"] = "struct"
	selector["2"] = "map"
	selector["3"] = "slice"
	selector["4"] = "for"
	selector["5"] = "interface"
	selector["6"] = "packages"
	selector["7"] = "channels"
	selector["8"] = "jsonHandler"
	fmt.Println("1) Run struct example")
	fmt.Println("2) Run map example")
	fmt.Println("3) Run slice example")
	fmt.Println("4) Run for loops example")
	fmt.Println("5) Run interfaces example")
	fmt.Println("6) Run packages example")
	fmt.Println("7) Run channels example")
	fmt.Println("8) Run json handler example")
	fmt.Scanf("%s", &option)

	return selector[option]
}
