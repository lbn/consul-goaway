all: run

build:
	gb build

run: build
	./bin/consul-goaway -consulAddr localhost:8500 -intervalS 60
