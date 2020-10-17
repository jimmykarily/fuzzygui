docker-builder:
	docker build -t fuzzygui-builder .

build-docker:
	docker run -it -v ${PWD}:/workspace -w /workspace fuzzygui-builder go build

build:
	go-bindata gui.glade
	go build
