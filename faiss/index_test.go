package faiss

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	index := &Index{
		Dimension:   128,
		Description: "Flat",
		MetricType:  L2{},
	}
	i, err := index.Create()
	if err != nil {
		t.Errorf("Index Create error ")
	}
	fmt.Println(i)
}
