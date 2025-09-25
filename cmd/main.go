package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const SW = 1200
const SH = 800

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

func (g *Grafo) Dijkstra(inicio, fim string) {

}

func main() {
	input_flag := flag.String("grafo", "grafo.json", "Arquivo de grafo para ser utilizado")
	ajuda_flag := flag.Bool("ajuda", false, "Imprime ajuda")
	help_flag := flag.Bool("help", false, "Imprime ajuda")
	inicio_flag := flag.String("inicio", "", "Estado inicial")
	fim_flag := flag.String("fim", "", "Estado final")
	flag.Parse()

	if *help_flag || *ajuda_flag {
		flag.Usage()
		return
	}

	if *inicio_flag == *fim_flag {
		fmt.Println("`-inicio` e `-fim` nao podem ser iguais")
		flag.Usage()
		return
	}

	if *inicio_flag == "" {
		fmt.Println("`-inicio` nao foi infomado")
		flag.Usage()
		return
	}

	if *fim_flag == "" {
		fmt.Println("`-fim` nao foi infomado")
		flag.Usage()
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

	cs := map[string]CircleNode{}
	edges := map[string][]Estado{}

	for i, v := range grafo_file.Vertices {
		x, y := (150 + i*50), 100
		cs[v.Nome] = CircleNode{
			Pos: rl.Vector2{
				X: rl.Clamp(float32(x)+rand.Float32()*SW, 150, SW-150),
				Y: rl.Clamp(float32(y)+rand.Float32()*SH, 150, SH-150),
			},
			Radius: 20,
		}
		for _, e := range v.Estados {
			edges[v.Nome] = append(edges[v.Nome], e)
		}
	}

	rl.InitWindow(SW, SH, "Dijkstra")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for cnome, c := range cs {
			for _, e := range edges[cnome] {
				rl.DrawLineV(
					c.Pos,
					cs[e.Nome].Pos,
					rl.DarkGray,
				)
			}

			rl.DrawCircleV(c.Pos, c.Radius, rl.Red)
			rl.DrawText(cnome, int32(c.Pos.X)-8, int32(c.Pos.Y)-8, 16, rl.Black)
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

type CircleNode struct {
	// Name string
	Pos    rl.Vector2
	Radius float32
}
