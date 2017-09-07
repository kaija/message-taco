
.PHONY: build test pack

all: build

build:
	go build

test:
	go test

docker_pack:
	docker build -t message-taco  -f Dockerfile .

docker_run:
	docker run -it --rm  message-taco
