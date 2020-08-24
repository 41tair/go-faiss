# Go-Faiss

[Faiss](https://github.com/facebookresearch/faiss) is a library for efficient similarity search and clustering of dense vectors.

Go-Faiss is go library use Faiss directly

	Faiss Version v1.6.3
## Install
### Build from source
#### Build Faiss
https://github.com/facebookresearch/faiss/blob/v1.6.3/INSTALL.md
#### Build C interface
https://github.com/facebookresearch/faiss/blob/v1.6.3/c_api/INSTALL.md

### Download from github

```
	git clone https://github.com/41tair/go-faiss
	mv go-faiss/lib/libfaiss_c.so /usr/local/lib/libfaiss_c.go
```

## Use in go

```
    import "github.com/41tair/go-faiss/faiss"
```


## Test

```
    go test
```
