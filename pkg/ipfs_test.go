package perm

import (
	"testing"
)

func TestAdd(t *testing.T) {
	err := Add("testdata")
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	err := Get("QmXHr4Am5BaifY5icf4ZCzXwNPrsk4RGPmZYv49p2DgttG", "QmXHr4Am5BaifY5icf4ZCzXwNPrsk4RGPmZYv49p2DgttG")
	if err != nil {
		t.Error(err)
	}
}
