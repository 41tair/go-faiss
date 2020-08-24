package faiss

import (
	"testing"
)

func TestInsertVectors(t *testing.T) {
	index := &Index{
		Dimension:   128,
		Description: "IDMap,Flat",
		MetricType:  L2{},
	}
	i, _ := index.Create()
	v := GenVectors(100, 128)
	ids := GenIDs(100)
	err := InsertVectors(v, i, 128, ids)
	if err != nil {
		t.Errorf("Insert error")
	}
}

func TestSearch(t *testing.T) {
	index := &Index{
		Dimension:   128,
		Description: "IDMap,Flat",
		MetricType:  L2{},
	}
	i, _ := index.Create()
	v := GenVectors(1000, 128)
	ids := GenIDs(1000)
	err := InsertVectors(v, i, 128, ids)
	if err != nil {
		t.Errorf("Insert error")
	}
	v2 := GenVectors(1, 128)
	resIDs := make([]int32, 10*1000*100)
	resDistances := make([]float32, 10*1000*100)
	SearchVectors(v2, i, 1000, 10, resIDs, resDistances)
}
