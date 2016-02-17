all: test

clean:
	go clean

make:
	go test

test: clean
	go build
	./ipfs-perm-add
