BINDING_DIR := ./gofaiss
gen:
	c-for-go -out ./ $(BINDING_DIR)/faiss.yml
	cd $(BINDING_DIR); \
#		go tool cgo -godefs types.go > types_gen.go; \
#		rm -rf types.go _obj
