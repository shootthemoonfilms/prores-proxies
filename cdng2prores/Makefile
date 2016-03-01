NAME    := prores-proxies
GO      := /opt/go
VERSION := v0.2
SUFFIX  := $(VERSION).

all: clean build

build:
	GOROOT=$(GO) CGO_ENABLED=0 GOOS=darwin  GOARCH=386   go build -o $(NAME).$(SUFFIX)x86.darwin
	GOROOT=$(GO) CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o $(NAME).$(SUFFIX)x64.darwin
	GOROOT=$(GO) CGO_ENABLED=0 GOOS=linux   GOARCH=386   go build -o $(NAME).$(SUFFIX)x86.linux
	GOROOT=$(GO) CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o $(NAME).$(SUFFIX)x64.linux
	GOROOT=$(GO) CGO_ENABLED=0 GOOS=windows GOARCH=386   go build -o $(NAME).$(SUFFIX)win32.exe
	GOROOT=$(GO) CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(NAME).$(SUFFIX)win64.exe

clean:
	rm -f *.darwin *.linux *.exe

