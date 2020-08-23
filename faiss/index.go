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

type IndexFactory interface {
	Create() (*C.FaissIndex, error)
	Turn() CIndex
	MetricMap() C.FaissMetricType
}

type MetricTyper interface {
	Name() string
}

type Index struct {
	Dimension   int
	Description string
	MetricType  MetricTyper
}

type CIndex struct {
	Dimension   C.int
	Description *C.char
	MetricType  C.FaissMetricType
}

func (i *Index) Create() (*C.FaissIndex, error) {
	ci := i.Turn()
	return C.CreateIndex(ci.Dimension, ci.Description, ci.MetricType), nil
}

func (i *Index) Turn() CIndex {
	return CIndex{
		Dimension:   C.int(i.Dimension),
		Description: C.CString(i.Description),
		MetricType:  i.MetricMap(),
	}
}

type L2 struct{}

func (m L2) Name() string {
	return "L2"
}

func (i *Index) MetricMap() C.FaissMetricType {
	mm := map[string]C.FaissMetricType{
		"L2": C.METRIC_L2,
	}
	if res, ok := mm[i.MetricType.Name()]; ok {
		return res
	}
	return C.METRIC_L2
}
