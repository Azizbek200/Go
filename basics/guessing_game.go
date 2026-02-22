package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//generate random number between 1 and 100
	target := random.Intn(100) + 1

	//welcome message
	fmt.Println("welcome to the game: ")

	var guess int
	for {
		fmt.Println("enter your input: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("congrats")
			break
		} else if guess < target {
			fmt.Println("too low")
		} else {
			fmt.Println("too high")
		}

	}

}
