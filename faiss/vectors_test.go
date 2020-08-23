package faiss

import (
	"testing"
)

func TestInsertVectors(t *testing.T) {
	index := &Index{
		Dimension:   128,
		Description: "Flat",
		MetricType:  L2{},
	}
	i, _ := index.Create()
	v := GenVectors(100, 128)
	err := InsertVectors(v, i, 128)
	if err != nil {
		t.Errorf("Insert error")
	}
}

func TestSearch(t *testing.T) {
	index := &Index{
		Dimension:   128,
		Description: "Flat",
		MetricType:  L2{},
	}
	i, _ := index.Create()
	v := GenVectors(100000, 128)
	err := InsertVectors(v, i, 128)
	if err != nil {
		t.Errorf("Insert error")
	}
	SearchVectors(v, i, 100000, 5)
}
