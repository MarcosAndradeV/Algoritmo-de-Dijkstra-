all: build/linux_dijkstra build/windows_dijkstra.exe

build/linux_dijkstra: cmd/main.go
	GOOS=linux GOARCH=amd64 go build -o build/linux_dijkstra cmd/main.go

build/windows_dijkstra.exe: cmd/main.go
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -o build/windows_dijkstra.exe cmd/main.go
