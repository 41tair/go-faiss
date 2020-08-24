package faiss

//#cgo CFLAGS: -I./include
//#cgo LDFLAGS: -L/usr/local/lib -lfaiss_c
//
//#include "faiss_cgo.h"
//#include "error_c.h"
//#include "index_io_c.h"
//#include "index_factory_c.h"
//#include "Index_c.h"
//#include "IndexFlat_c.h"
//#include "AutoTune_c.h"
//#include "clone_index_c.h"
import "C"

import (
	"math/rand"
	"time"
	"unsafe"
)

type Result struct {
	Distance float32
	ID       int32
}

type Vectors []float32

func GenVectors(dataSize int, dimension int) Vectors {
	rand.Seed(time.Now().UnixNano())
	vectors := make([]float32, dataSize*dimension)
	for i := 0; i < dataSize; i++ {
		for j := 0; j < dimension; j++ {
			d := rand.Float32()
			vectors[dimension*i+j] = d
			vectors[dimension*i] += float32(i) / float32(1000)
		}
	}
	return vectors
}

func GenIDs(dataSize int) []int32 {
	rand.Seed(time.Now().UnixNano())
	ids := make([]int32, dataSize)
	for i := 0; i < dataSize; i++ {
		ids[i] = rand.Int31()
	}
	return ids
}
func InsertVectors(v Vectors, index *C.FaissIndex, dimension int, ids []int32) error {
	C.Insert((*C.float)(unsafe.Pointer(&v[0])), index, C.int(dimension), (*C.long)(unsafe.Pointer(&ids[0])))
	return nil
}

func SearchVectors(v Vectors, index *C.FaissIndex, nq int, topk int, resIDs []int32, resDistances []float32) []Result {
	res := []Result{}
	C.Search((*C.float)(unsafe.Pointer(&v[0])), index, C.int(nq), C.int(topk),
		(*C.long)(unsafe.Pointer(&resIDs[0])), (*C.float)(unsafe.Pointer(&resDistances[0])))
	for i := 0; i < topk; i++ {
		res = append(res, Result{
			Distance: resDistances[i],
			ID:       resIDs[i],
		})
	}
	return res
}
