package console

import (
	"fmt"
	"strings"
)

// Init is to create an initialized text
func Init() {

	fmt.Printf("\n%s\n", strings.Repeat("*", 80))
	fmt.Printf("Welcome to echGo! Here you will find the most important information directly\nin the console. For more information visit https://github.com/echgo/echgo.\n")
	fmt.Printf("%s\n\n", strings.Repeat("*", 80))

}
