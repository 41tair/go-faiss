#include "faiss_cgo.h"
#include <stdio.h>
#include <stdlib.h>

#include <time.h>

#include "include/error_c.h"
#include "include/index_io_c.h"
#include "include/index_factory_c.h"
#include "include/Index_c.h"
#include "include/IndexFlat_c.h"
#include "include/AutoTune_c.h"
#include "include/clone_index_c.h"

#define FAISS_TRY(C)                                       \
    {                                                      \
        if (C) {                                           \
            fprintf(stderr, "%s", faiss_get_last_error()); \
            exit(-1);                                      \
        }                                                  \
    }


FaissIndex* CreateIndex(int dimension, const char* description, FaissMetricType metric_type) {
  FaissIndex* index = NULL;
  FAISS_TRY(faiss_index_factory(&index, dimension, description, metric_type));
  return index;
}

void Insert(float* vectors, FaissIndex* index, int num, long* ids) {
  printf("Start insert \n");
  FAISS_TRY(faiss_Index_add_with_ids(index, num, vectors, ids));
  //  printf("%d", faiss_Index_add(index,num,vectors));
  printf("Insert finish \n");
  return;
}

void Search(float* vectors, FaissIndex* index, int nq, int topk, long* ids, float* distances) {
  printf("Start search");
  FAISS_TRY(faiss_Index_search(index, nq, vectors, topk, distances, ids));
}
