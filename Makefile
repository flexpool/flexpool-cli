.PHONY: make run

make:
	go build

install:
	go install

uninstall:
	go clean -i