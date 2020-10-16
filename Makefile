docker-builder:
	docker build -t fuzzygui-builder .

build:
	docker run -it -v ${PWD}:/workspace -w /workspace fuzzygui-builder go build
