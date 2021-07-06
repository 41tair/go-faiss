BINDING_DIR := ./gofaiss
gen:
	c-for-go -ccincl -out ./ $(BINDING_DIR)/faiss.yml
	cd ./faiss; \
	# go tool cgo -godefs types.go > types_gen.go; \
	# rm -rf types.go _obj
