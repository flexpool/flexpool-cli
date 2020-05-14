package utils

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// Ask4confirm asks y/n
func Ask4confirm() bool {
	var yes bool

	fmt.Print("(y/n): ")
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}

	if char == 'y' {
		yes = true
	}

	fmt.Println(string(char))
	return yes
}
