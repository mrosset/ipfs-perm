
all: test

clean:
	go clean

make:
	go test

test: clean
	make -C ipfs-perm
