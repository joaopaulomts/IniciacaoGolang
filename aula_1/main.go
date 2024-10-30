package main

import "fmt"

func main() {
	var media float32
	var name string

	fmt.Println("Qual seu nome? ")
	fmt.Scanln(&name)

	for i := 0; i < 4; i++ {
		var nota float32
		fmt.Printf("Qual a sua nota %d: ", i+1)
		fmt.Scanln(&nota)

		media += nota
	}

	fmt.Printf("Seu nome é %s e sua média é %.2f", name, media/float32(4))

}
