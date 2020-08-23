#include "include/error_c.h"
#include "include/index_io_c.h"
#include "include/index_factory_c.h"
#include "include/Index_c.h"
#include "include/IndexFlat_c.h"
#include "include/AutoTune_c.h"
#include "include/clone_index_c.h"

FaissIndex* CreateIndex(int dimension, const char* description, FaissMetricType metric_type);

void Insert(float* vectors, FaissIndex* index, int num);

void Search(float* vectors, FaissIndex* index, int nq, int topk);
