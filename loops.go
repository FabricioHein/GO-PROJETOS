package main

import "fmt"

func main() {
    // For Tradicional
    for i := 0; i < 5; i++ {
        fmt.Println("For Tradicional:", i)
    }

    // For como While
    contador := 0
    for contador < 5 {
        fmt.Println("For como While:", contador)
        contador++
    }

    // For com Range
    numeros := []int{1, 2, 3, 4, 5}
    for indice, valor := range numeros {
        fmt.Printf("For com Range - Índice: %d, Valor: %d\n", indice, valor)
    }
}
