package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Estado struct {
	Nome string `json:"nome"`
	Peso int    `json:"peso"`
}

type Vertice struct {
	Nome    string   `json:"nome"`
	Estados []Estado `json:"estados"`
}

type Grafo struct {
	Vertices []Vertice `json:"vertices"`
}

func (g *Grafo) TemEstado(e string) bool {
	for _, v := range g.Vertices {
		if v.Nome == e {
			return true
		}
	}
	return false
}


func (g *Grafo) Dijkstra(inicio, fim string) bool {
	fmt.Println("Dijkstra não implementado")
	os.Exit(1)
	return false
}

func main() {
	input_flag := flag.String("grafo", "grafo.json", "Arquivo de grafo para ser utilizado")
	ajuda_flag := flag.Bool("ajuda", false, "Imprime ajuda")
	help_flag := flag.Bool("help", false, "Imprime ajuda")
	inicio_flag := flag.String("inicio", "", "Estado inicial")
	fim_flag := flag.String("fim", "", "Estado final")
	flag.Parse()

	if (*help_flag || *ajuda_flag) {
		flag.Usage()
		return
	}

	if *inicio_flag == *fim_flag {
		fmt.Println("`-inicio` e `-fim` nao podem ser iguais")
		flag.Usage();
		return
	}

	if *inicio_flag == "" {
		fmt.Println("`-inicio` nao foi infomado")
		flag.Usage();
		return
	}

	if *fim_flag == "" {
		fmt.Println("`-fim` nao foi infomado")
		flag.Usage();
		return
	}

	data, err := os.ReadFile(*input_flag)
	if err != nil {
		fmt.Printf("Problema ao ler o arquivo %s\n", *input_flag)
		fmt.Println(err)
		return
	}
	var grafo_file Grafo
	err = json.Unmarshal(data, &grafo_file)
	if err != nil {
		fmt.Printf("Problema ao ler o arquivo %s\n", *input_flag)
		fmt.Println(err)
		return
	}

	if !grafo_file.TemEstado(*inicio_flag) {
		fmt.Printf("O grafo `%s` não posui estado `%s`\n", *input_flag, *inicio_flag)
	}

	if !grafo_file.TemEstado(*fim_flag) {
		fmt.Printf("O grafo `%s` não posui estado `%s`\n", *input_flag, *fim_flag)
	}

	fmt.Println(grafo_file)

	grafo_file.Dijkstra(*inicio_flag, *fim_flag)
}
