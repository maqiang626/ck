export tag=v1.0
root:
	export ROOT=github.com/maqiang626/ck

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t maqiang626/httpserver:${tag} .

push: release
	echo "pushing maqiang626/httpserver"
	docker push maqiang626/httpserver:${tag}
