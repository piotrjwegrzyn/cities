all: build run

.PHONY: build
build:
	go build -o bin/ .

.PHONY: run
run:
	bin/cities -port 8081

.PHONY: clean
clean:
	rm -r bin/