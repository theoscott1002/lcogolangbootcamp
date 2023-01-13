package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var name string
	var userRating string

	//Front end
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name")
	name, _ = reader.ReadString('\n')

	//take rating from user and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our Dosa center between 1 and 5: ")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Back end
	fmt.Println("%v, %v", name, mynumRating)
}
