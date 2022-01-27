package main

import (
	"fmt"
)

func main() {
	fmt.Println(stage())
	// fmt.Println("   1  |   2  |   3  ")
	// fmt.Println("      |      |      ")
	// fmt.Println("      |      |      ")
	// fmt.Println("--------------------")
	// fmt.Println("   4  |   5  |   6  ")
	// fmt.Println("      |      |      ")
	// fmt.Println("      |      |      ")
	// fmt.Println("--------------------")
	// fmt.Println("   7  |   8  |   9  ")
	// fmt.Println("      |      |      ")
	// fmt.Println("      |      |      ")
}

func stage() string {
	stage := "|   1  |   2  |   3  |\n|      |      |      |\n|      |      |      |\n|--------------------|\n|   4  |   5  |   6  |\n|      |      |      |\n|      |      |      |\n|--------------------|\n|   7  |   8  |   9  |\n|      |      |      |\n|      |      |      |"
	return stage
}
