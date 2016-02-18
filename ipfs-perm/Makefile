
all: test

clean:
	go clean

make:
	go test

test: clean
	go build
	./ipfs-perm add testdata
	./ipfs-perm get QmXHr4Am5BaifY5icf4ZCzXwNPrsk4RGPmZYv49p2DgttG
