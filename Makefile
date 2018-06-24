install:
	dep ensure

build:
	go build

clean:
	rm ./mytime

dev:
	go build
	./mytime -e development