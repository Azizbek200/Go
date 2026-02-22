package basics

import "fmt"

func main() {

	//simple iteration over a range

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	//iterate over collection

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, value %d\n", index, value)
	}
}
