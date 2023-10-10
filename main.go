package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Defina uma função de manipulador para a rota raiz "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Escreva a resposta no ResponseWriter
        fmt.Fprint(w, "Bem-vindo à minha API Go!")
    })

    // Inicie o servidor na porta 8080
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
    }
}
