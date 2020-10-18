all: deps build

build:
	go-bindata gui.glade
	go build -o release/fuzzygui

deps:
	go install github.com/go-bindata/go-bindata/...

docker-builder:
	docker build -t fuzzygui-builder .

build-docker:
	docker run -it -v ${PWD}:/workspace -w /workspace fuzzygui-builder make all
