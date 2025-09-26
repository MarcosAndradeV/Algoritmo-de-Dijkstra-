all: build/linux_dijkstra build/windows_dijkstra.exe

VERSION=ce4f07a84931cf82bb383406d40124194fd377f8

LDFLAGS=-ldflags "-w -s -X main.version=${VERSION}"
LDFLAGSWIN=-ldflags "-w -s -X main.version=${VERSION} -H=windowsgui"

build/linux_dijkstra: cmd/main.go
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/linux_dijkstra cmd/main.go

build/windows_dijkstra.exe: cmd/main.go
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build ${LDFLAGSWIN} -o build/windows_dijkstra.exe cmd/main.go
