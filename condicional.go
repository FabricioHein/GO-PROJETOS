package main

import "fmt"

func main() {
    nota := 75
	if nota >= 90 {
		fmt.Println("Aprovado com distinção.")
	} else if nota >= 60 {
		fmt.Println("Aprovado.")
	} else {
		fmt.Println("Reprovado.")
}
}
