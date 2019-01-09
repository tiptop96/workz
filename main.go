package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out, quit := CreateWorkers(10)

	i := 0

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')

		switch text {
		case "quit\n":
			quit()
			return
		default:
			Jobs <- Job{i, text}
			fmt.Print(<-out)
		}
		i++
	}
}
