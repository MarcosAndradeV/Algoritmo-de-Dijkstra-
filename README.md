# Como usar

## Build

- Instalar [Go](https://go.dev/) e Make

```console
make
```

ou

- Instalar [Go](https://go.dev/)

### Para Windows
```console
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build ${LDFLAGSWIN} -o build/windows_dijkstra.exe cmd/main.go
```

### Para Linux
```console
GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/linux_dijkstra cmd/main.go
```

## Como Usar o Programa

```console
$ linux_dijkstra -help
Usage of ./build/linux_dijkstra:
  -fim string
    	fim do grafo (default "E")
  -grafo string
    	Arquivo do grafo em json (default "grafo.json")
  -help
    	Ajuda
  -inicio string
    	inicio do grafo (default "A")
```

### Controles:

- S (Avançar): Pressione a tecla "S" para avançar um passo na execução do algoritmo de Dijkstra. Isso permite que você veja exatamente o que o algoritmo faz a cada iteração.

- R (Resetar Posições): Pressione a tecla "R" para randomizar a posição dos círculos (nós) na tela, caso eles fiquem sobrepostos ou mal distribuídos.

- Arrastar os Nós: Você pode clicar com o botão esquerdo do mouse sobre qualquer círculo e arrastá-lo para reposicioná-lo na tela. Isso ajuda a organizar a visualização do grafo da maneira que preferir.
