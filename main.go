package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type GrafoDoArquivo struct {
	A struct {
		B int `json:"B"`
		C int `json:"C"`
	} `json:"A"`
	B struct {
		A int `json:"A"`
		C int `json:"C"`
		D int `json:"D"`
	} `json:"B"`
	C struct {
		A int `json:"A"`
		B int `json:"B"`
		D int `json:"D"`
	} `json:"C"`
	D struct {
		B int `json:"B"`
		C int `json:"C"`
		E int `json:"E"`
	} `json:"D"`
	E struct {
		D int `json:"D"`
		F int `json:"F"`
	} `json:"E"`
	F struct {
		E int `json:"E"`
	} `json:"F"`
}

func main() {
	data, err := os.ReadFile("grafo.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var grafo_file GrafoDoArquivo
	err = json.Unmarshal(data, &grafo_file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(grafo_file)
}
