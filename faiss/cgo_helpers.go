// WARNING: This file has automatically been generated on Tue, 06 Jul 2021 19:41:13 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package faiss

/*
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files
#include "faiss_c.h"
#include "Index_c.h"
#include "AutoTune_c.h"
#include "clone_index_c.h"
#include "Clustering_c.h"
#include "error_c.h"
#include "index_factory_c.h"
#include "IndexFlat_c.h"
#include "index_io_c.h"
#include "IndexIVF_c.h"
#include "IndexIVFFlat_c.h"
#include "IndexLSH_c.h"
#include "IndexPreTransform_c.h"
#include "IndexShards_c.h"
#include "MetaIndexes_c.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(h.Data), C.int(h.Len))
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// copyPDoubleBytes copies the data from Go slice as *C.double.
func copyPDoubleBytes(slice *sliceHeader) (*C.double, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: slice.Data,
		Len:  int(sizeOfDoubleValue) * slice.Len,
		Cap:  int(sizeOfDoubleValue) * slice.Len,
	}))))
	allocs.Add(mem0)

	return (*C.double)(mem0), allocs
}

// allocDoubleMemory allocates memory for type C.double in C.
// The caller is responsible for freeing the this memory via C.free.
func allocDoubleMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfDoubleValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfDoubleValue = unsafe.Sizeof([1]C.double{})

// safeString ensures that the string is NULL-terminated, a NULL-terminated copy is created otherwise.
func safeString(str string) string {
	if len(str) > 0 && str[len(str)-1] != '\x00' {
		str = str + "\x00"
	} else if len(str) == 0 {
		str = "\x00"
	}
	return str
}

// unpackPCharString copies the data from Go string as *C.char.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	str = safeString(str)
	mem0 := unsafe.Pointer(C.CString(str))
	allocs.Add(mem0)
	return (*C.char)(mem0), allocs
}

// copyPFloatBytes copies the data from Go slice as *C.float.
func copyPFloatBytes(slice *sliceHeader) (*C.float, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: slice.Data,
		Len:  int(sizeOfFloatValue) * slice.Len,
		Cap:  int(sizeOfFloatValue) * slice.Len,
	}))))
	allocs.Add(mem0)

	return (*C.float)(mem0), allocs
}

// allocFloatMemory allocates memory for type C.float in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFloatMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFloatValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFloatValue = unsafe.Sizeof([1]C.float{})

// copyPIdxBytes copies the data from Go slice as *C.idx_t.
func copyPIdxBytes(slice *sliceHeader) (*C.idx_t, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: slice.Data,
		Len:  int(sizeOfIdxValue) * slice.Len,
		Cap:  int(sizeOfIdxValue) * slice.Len,
	}))))
	allocs.Add(mem0)

	return (*C.idx_t)(mem0), allocs
}

// allocIdxMemory allocates memory for type C.idx_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocIdxMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfIdxValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfIdxValue = unsafe.Sizeof([1]C.idx_t{})