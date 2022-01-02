ROOT_DIR := $(shell pwd)

.PHONY: all compile

all: compile

docker-build: # Builds the docker image
	docker build -f docker/seekret-public-protoc.dockerfile -t gcr.io/seekret/seekret-public-protoc .

docker-push:
	docker push gcr.io/seekret/seekret-public-protoc

compile: # Compiles the files
	docker run --rm -w /src -it -v ${ROOT_DIR}:/src gcr.io/seekret/seekret-public-protoc:latest /bin/bash -c "python3 ./generate.py --all --node-modules=/build/node_modules"
