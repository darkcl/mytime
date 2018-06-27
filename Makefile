install:
	dep ensure

build:
	go build

clean:
	rm ./mytime

dev:
	go build
	./mytime -e development

doc:
	aglio -i ./docs/spec.apib -o ./docs/index.html
	open ./docs/index.html