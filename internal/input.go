package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadInt(promt string) int {

	for {
		fmt.Print(promt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}
		return value
	}
}

func ReadString(promt string) string {
	fmt.Print(promt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
