package main

import (
	"fmt"

	"github.com/isaadabbasi/go_crud/initializers"
)

func main() {
	initializers.InitServer("3000", func(port string) {
		fmt.Println("Server Running at port " + port)
	})
}
