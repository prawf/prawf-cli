build:
	go mod tidy
	go build -o prawf main.go

test-1:
	./prawf -h
	./prawf run get
	./prawf init
	./prawf run get
	./prawf init

test-2:
	./prawf -h
	./prawf init
	./prawf run get
	./prawf init

test-3:
	./prawf -h
	./prawf init
	./prawf run get
	./prawf init -p . -n anotherprawf

test-4:
	./prawf -h
	./prawf init -p . -n anotherprawf
	./prawf run get
	./prawf init -p . -n anotherprawf

test: build test-1 test-2 test-3 test-4
