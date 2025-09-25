package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Estado struct {
	Nome string `json:"nome"`
	Peso int `json:"peso"`
}

type Vertice struct {
	Nome string `json:"nome"`
	Estados []Estado `json:"estados"`
}

type Grafo struct {
	Vertices []Vertice `json:"vertices"`
}

func main() {
	data, err := os.ReadFile("grafo.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var grafo_file Grafo
	err = json.Unmarshal(data, &grafo_file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(grafo_file)
}
