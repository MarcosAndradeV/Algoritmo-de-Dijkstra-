package main

type Pair struct {
	a, b int
}

type Grafo struct {
	V int
	Adj [][]Pair
}

func (g *Grafo) AddAresta(v1 int, v2 int, custo int) {
	g.Adj[v1] = append(&g.Adj[v1], Pair{v2, custo});
}
