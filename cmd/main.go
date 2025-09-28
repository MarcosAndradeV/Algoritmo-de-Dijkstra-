package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	JANELA_LARGURA   = 800
	JANELA_ALTURA    = 600
	RAIO             = 20
	DISTANCIA_MINIMA = 10
)

func main() {

	f := flag.String("grafo", "grafo.json", "Arquivo do grafo em json")

	data, err := os.ReadFile("grafo.json")
	if err != nil {
		fmt.Printf("WARN: não foi possivel ler o arquivo %s: %s\n", *f, err)
		data = []byte(GRAFO_JSON)
	}
	g := Grafo{}
	err = json.Unmarshal(data, &g)
	if err != nil {
		fmt.Printf("ERROR: não foi possivel ler o arquivo %s: %s\n", *f, err)
		os.Exit(1)
	}

	rl.InitWindow(JANELA_LARGURA, JANELA_ALTURA, "Dijkstra")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	nos := gerarCirculos(g.Nos)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		if rl.IsKeyPressed(rl.KeyR) {
			nos = gerarCirculos(g.Nos)
		}
		for _, a := range g.Arrestas {
			a.Desenhar(nos)
		}
		for _, c := range nos {
			c.Desenhar()
		}
		rl.EndDrawing()
	}
}

type Grafo struct {
	Nos      []No     `json:"nos"`
	Arrestas []Aresta `json:"arrestas"`
}

func (a *Aresta) Desenhar(nos []Circulo) {
	c1 := AcharCirculo(nos, a.N1)
	c2 := AcharCirculo(nos, a.N2)

	rl.DrawLineEx(c1.Pos, c2.Pos, 5, rl.DarkGray)

	midX := (c1.Pos.X + c2.Pos.X) / 2
	midY := (c1.Pos.Y + c2.Pos.Y) / 2

	dx := c2.Pos.X - c1.Pos.X
	dy := c2.Pos.Y - c1.Pos.Y

	len := float32(math.Sqrt(float64(dx*dx + dy*dy)))
	nx := -dy / len
	ny := dx / len

	textX := midX + nx*15
	textY := midY + ny*15

	text := fmt.Sprint(a.Peso)
	textWidth := rl.MeasureText(text, 16)
	rl.DrawText(
		text,
		int32(textX)-int32(textWidth)/2,
		int32(textY)-8,
		16,
		rl.Black,
	)
}

func AcharCirculo(cs []Circulo, n string) Circulo {
	for _, c := range cs {
		if c.Nome == n {
			return c
		}
	}
	return Circulo{}
}

type No struct {
	Nome string `json:"nome"`
}

type Aresta struct {
	N1   string `json:"n1"`
	N2   string `json:"n2"`
	Peso int    `json:"peso"`
}

type Circulo struct {
	Nome string
	Raio float32
	Pos  rl.Vector2
}

func (c *Circulo) Desenhar() {
	rl.DrawCircleV(c.Pos, c.Raio, rl.Red)
	rl.DrawText(c.Nome, int32(c.Pos.X-(c.Raio/4)), int32(c.Pos.Y-(c.Raio/4)-3), int32(c.Raio), rl.Black)
}

func distancia(a, b rl.Vector2) float32 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func gerarCirculos(nos []No) []Circulo {
	circulos := make([]Circulo, 0, len(nos))
	i := 0
tentativa:
	for len(circulos) < len(nos) {
		pos := rl.NewVector2(
			rand.Float32()*(float32(JANELA_LARGURA)-2*RAIO)+RAIO,
			rand.Float32()*(float32(JANELA_ALTURA)-2*RAIO)+RAIO,
		)

		for _, c := range circulos {
			if distancia(pos, c.Pos) < c.Raio+RAIO+DISTANCIA_MINIMA {
				continue tentativa
			}
		}

		circulos = append(circulos, Circulo{Nome: nos[i].Nome, Raio: RAIO, Pos: pos})
		i += 1
	}

	return circulos
}

const GRAFO_JSON = `
{
  "nos" : [
    {
      "nome": "A"
    },
    {
      "nome": "B"
    }
  ],
  "arrestas" : [
    {
      "n1": "A",
      "n2": "B",
      "peso": 5
    }
  ]
}
`
