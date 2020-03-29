package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter your score out of 100:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	// To strip of the newlines, tabs, and regular spaces
	input = strings.TrimSpace(input)
	fmt.Println("Input value :", input)
	if err != nil {
		log.Fatal(err)
	}
	score, _ := strconv.Atoi(input)
	if score >= 80 {
		fmt.Println("You qualified in distinction!")
	} else if score >= 60 {
		fmt.Println("You qualified in Ist class .")
	} else {
		fmt.Println("Sorry fail!")
	}
}
