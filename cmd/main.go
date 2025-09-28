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
	JANELA_LARGURA      = 800
	JANELA_ALTURA       = 600
	RAIO                = 20
	TAMANHO_FONTE_NO    = 16
	TAMANHO_FONTE_TEXTO = 26
	DISTANCIA_MINIMA    = 20
)

type Dijkstra struct {
	Inicio, Fim, NoAtual string
	Distancias           map[string]int
	Visitados            map[string]bool
	Arestas              map[string][]Aresta
	Caminho              map[string]string
	Terminado            bool
}

func NewDijkstra(g *Grafo, inicio, fim string) Dijkstra {
	d := Dijkstra{
		Inicio:     inicio,
		Fim:        fim,
		NoAtual:    inicio,
		Distancias: make(map[string]int),
		Visitados:  make(map[string]bool),
		Arestas:    make(map[string][]Aresta),
		Caminho:    make(map[string]string),
	}

	for _, no := range g.Nos {
		d.Distancias[no.Nome] = math.MaxInt
	}
	d.Distancias[inicio] = 0

	for _, a := range g.Arestas {
		d.Arestas[a.N1] = append(d.Arestas[a.N1], a)
		d.Arestas[a.N2] = append(d.Arestas[a.N2], a)
	}

	return d
}

func (d *Dijkstra) Avancar(g *Grafo) {
	if d.Terminado {
		return
	}

	distanciaMenor := math.MaxInt
	proximo := ""

	for no, dist := range d.Distancias {
		if !d.Visitados[no] && dist < distanciaMenor {
			distanciaMenor = dist
			proximo = no
		}
	}

	if distanciaMenor == math.MaxInt {
		d.Terminado = true
		return
	}

	d.NoAtual = proximo

	if d.NoAtual == d.Fim {
		d.Terminado = true
		d.Visitados[d.NoAtual] = true
		return
	}

	d.Visitados[d.NoAtual] = true

	for _, a := range d.Arestas[d.NoAtual] {
		vizinho := a.N1
		if vizinho == d.NoAtual {
			vizinho = a.N2
		}

		if d.Visitados[vizinho] {
			continue
		}

		distancia := d.Distancias[d.NoAtual] + a.Peso

		if distancia < d.Distancias[vizinho] {
			d.Distancias[vizinho] = distancia
			d.Caminho[vizinho] = d.NoAtual
		}
	}
}

func main() {

	f := flag.String("grafo", "grafo.json", "Arquivo do grafo em json")
	inicio := flag.String("inicio", "A", "inicio do grafo")
	fim := flag.String("fim", "E", "fim do grafo")
	help := flag.Bool("help", false, "Ajuda")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	data, err := os.ReadFile(*f)
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
	var selected int
	var isSelected bool

	d := NewDijkstra(&g, *inicio, *fim)

	for !rl.WindowShouldClose() {
		if rl.IsKeyReleased(rl.KeyS) && !d.Terminado {
			d.Avancar(&g)
		}
		if rl.IsKeyPressed(rl.KeyR) {
			nos = gerarCirculos(g.Nos)
		}
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			mouse_pos := rl.GetMousePosition()
			for i, c := range nos {
				if rl.CheckCollisionPointCircle(mouse_pos, c.Pos, c.Raio) {
					selected = i
					isSelected = true
					break
				}
			}
		}
		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			isSelected = false
		}
		if isSelected {
			mouse_pos := rl.GetMousePosition()
			nos[selected].Pos = mouse_pos
		}
		for no, vis := range d.Visitados {
			if vis {
				AcharCirculo(nos, no).Cor = rl.Green
			}
		}
		if d.Terminado {
			no := d.Fim
			for {
				cno, ok := d.Caminho[no]
				if !ok {
					break
				}
				AcharCirculo(nos, no).Cor = rl.Gold
				AcharCirculo(nos, cno).Cor = rl.Gold
				no = cno
			}
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(
			fmt.Sprintf("Inicio: %s | Atual: %s | Fim: %s ", d.Inicio, d.NoAtual, d.Fim),
			10,
			10,
			TAMANHO_FONTE_TEXTO,
			rl.Black,
		)
		for _, a := range g.Arestas {
			a.Desenhar(nos)
		}
		for _, c := range nos {
			c.Desenhar()
		}
		rl.EndDrawing()
	}
}

type Grafo struct {
	Nos     []No     `json:"nos"`
	Arestas []Aresta `json:"arestas"`
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

func AcharCirculo(cs []Circulo, n string) *Circulo {
	for i := range cs {
		c := &cs[i]
		if c.Nome == n {
			return c
		}
	}
	return nil
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
	Cor  rl.Color
	Raio float32
	Pos  rl.Vector2
}

func (c *Circulo) Desenhar() {
	rl.DrawCircleV(c.Pos, c.Raio, c.Cor)
	rl.DrawText(c.Nome, int32(c.Pos.X-(c.Raio/4)), int32(c.Pos.Y-(c.Raio/4)-3), TAMANHO_FONTE_NO, rl.Black)
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
			if rl.Vector2Distance(pos, c.Pos) < c.Raio+RAIO+DISTANCIA_MINIMA {
				continue tentativa
			}
		}

		circulos = append(circulos, Circulo{Nome: nos[i].Nome, Cor: rl.Red, Raio: RAIO, Pos: pos})
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
    },
    {
      "nome": "C"
    },
    {
      "nome": "D"
    },
    {
      "nome": "E"
    }
  ],
  "arestas" : [
    {
      "n1": "A",
      "n2": "B",
      "peso": 5
    },
    {
      "n1": "A",
      "n2": "C",
      "peso": 10
    },
    {
      "n1": "B",
      "n2": "D",
      "peso": 7
    },
    {
      "n1": "D",
      "n2": "A",
      "peso": 3
    },
    {
      "n1": "D",
      "n2": "E",
      "peso": 2
    },
    {
      "n1": "B",
      "n2": "E",
      "peso": 6
    }
  ]
}
`
