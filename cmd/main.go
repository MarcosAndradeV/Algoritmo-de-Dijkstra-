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


var inicio_flag, fim_flag string

func main() {
	input_flag := flag.String("grafo", "grafo.json", "Arquivo de grafo para ser utilizado")
	ajuda_flag := flag.Bool("ajuda", false, "Imprime ajuda")
	help_flag := flag.Bool("help", false, "Imprime ajuda")
	flag.StringVar(&inicio_flag, "inicio", "", "Estado inicial")
	flag.StringVar(&fim_flag, "fim", "", "Estado final")
	flag.Parse()

	if *help_flag || *ajuda_flag {
		flag.Usage()
		return
	}

	if inicio_flag == fim_flag {
		fmt.Println("`-inicio` e `-fim` nao podem ser iguais")
		flag.Usage()
		return
	}

	if inicio_flag == "" {
		fmt.Println("`-inicio` nao foi infomado")
		flag.Usage()
		return
	}

	if fim_flag == "" {
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

	if !grafo_file.TemEstado(inicio_flag) {
		fmt.Printf("O grafo `%s` não posui estado `%s`\n", *input_flag, inicio_flag)
	}

	if !grafo_file.TemEstado(fim_flag) {
		fmt.Printf("O grafo `%s` não posui estado `%s`\n", *input_flag, fim_flag)
	}

	cs := map[string]*CircleNode{}
	edges := map[string][]Estado{}

	shuffleNodes(grafo_file, cs, edges)

	rl.InitWindow(SW, SH, "Dijkstra")

	var selected_c string

	for !rl.WindowShouldClose() {
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) && selected_c != "" {
		 	cs[selected_c].Pos = rl.GetMousePosition()
			selected_c = ""
		}
		if rl.IsKeyPressed(rl.KeyR) {
			shuffleNodes(grafo_file, cs, edges)
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for cnome, c := range cs {
			cpos := c.Pos
			ccolor := rl.Red
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				mouse_pos := rl.GetMousePosition()
				if rl.CheckCollisionPointCircle(mouse_pos, c.Pos, c.Radius) {
					selected_c = cnome
				}
			}
			if selected_c == cnome {
				cpos = rl.GetMousePosition()
				ccolor = rl.Gold
			}
			for _, e := range edges[cnome] {
				rl.DrawLineV(
					cpos,
					cs[e.Nome].Pos,
					rl.DarkGray,
				)
				rl.DrawText(
					fmt.Sprintf("%d", e.Peso),
					int32(cpos.X + cs[e.Nome].Pos.X)/2,
					int32(cpos.Y + cs[e.Nome].Pos.Y)/2,
					26,
					rl.Black,
				)
			}
			rl.DrawCircleV(cpos, c.Radius, ccolor)
			rl.DrawText(cnome, int32(cpos.X)-8, int32(cpos.Y)-8, 16, rl.Black)
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func shuffleNodes(grafo_file Grafo, cs map[string]*CircleNode, edges map[string][]Estado) {
	for i, v := range grafo_file.Vertices {
		x, y := (150 + i*50), 100
		c := new(CircleNode)
		*c = CircleNode{
			Pos: rl.Vector2{
				X: rl.Clamp(float32(x)+rand.Float32()*SW, 150, SW-150),
				Y: rl.Clamp(float32(y)+rand.Float32()*SH, 150, SH-150),
			},
			Radius: 20,
		}
		cs[v.Nome] = c
		for _, e := range v.Estados {
			edges[v.Nome] = append(edges[v.Nome], e)
		}
	}
}

type CircleNode struct {
	// Name string
	Pos    rl.Vector2
	Radius float32
}
