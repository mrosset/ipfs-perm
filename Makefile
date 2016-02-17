
all: test

clean:
	go clean

make:
	go test

test: clean
	go build
	./ipfs-perm add testdata
	./ipfs-perm get QmfRz3rYH4ZBuksyo8acRYnVhMcPQsEVtPxo6VCQRMcepb
