package main // Package declaration, we use main package to create executable file on main function
import "fmt" // Importing fmt package, fmt package is used to print the output to the console
// we have to import the package to use the function in that package golang.org/pkg

func main() { // this function will turn into executable file
	fmt.Println("Hello, World!")      // Println is a function in fmt package
	var card string = "Ace of Spades" // Declaring a variable of type string, go statically typed language you have to declare the type of variable
	fmt.Println(card)

	//or you can use := to declare and assign the value
	card2 := "Five of Diamonds"
	fmt.Println(card2)

	//use := to declare and assign the value, use = to assign the value to the variable that is already declared
	card2 = "Six of Diamonds"
	fmt.Println(card2)

	//function call
	card3 := newCard()
	fmt.Println(card3)

	//slice
	//slice is a dynamic array, you can add or remove the element from the slice

	cards := []string{"Ace of Diamonds", newCard()} // slice of type string
	cards = append(cards, "Six of Spades")          // append function to add the element to the slice
	// it will return the new slice with the new element

	//for loop
	for i, card := range cards { // range function to iterate over the slice
		fmt.Println(i, card)
	}
}

func newCard() string { // function declaration, return type is string
	return "Seven of Diamonds"
}
