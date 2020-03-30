package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
func main() {
	var srnd = random(1, 100)
	fmt.Println("Guess, What was the system generated a random number that falls between 1 to 100")
	fmt.Println("You can discover in 10 guesses.")
	var x int = 1
	for x <= 10 {
		fmt.Println("Guess ", x, "/10 what was system generated random number:")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		rv, _ := strconv.Atoi(strings.TrimSpace(input))
		if srnd == rv {
			fmt.Println("Good job! You guessed it!", x, "/10")
			break
		} else if srnd != rv && x == 10 {
			fmt.Println("Sorry. You didn't guess my number. It was:", srnd)
		} else if srnd > rv {
			fmt.Println("Oops. Your guess was LOW")
		} else if srnd < rv {
			fmt.Println("Oops. Your guess was HIGH")
		}
		x = x + 1

	}
}
