package main

import (
	"fmt"
	"os"
)

func main() {

	argCodigoUsuario := os.Args[1]
	switch argCodigoUsuario {
	case "1234567":
		fmt.Println("ID: FLUNA")
		fmt.Println("User: Fernando Luna")
	case "7654321":
		fmt.Println("ID: WCRESPO")
		fmt.Println("User: Wilson Crespo")
	default:
		fmt.Println("Unknown")
	}

}
