run this in the package directory instead of ```go build```

```go
gobuildall
```

this will create a folder `./build/` with all the binary compiled files for all the allowed permutations of GOOS and GOARCH as per this doc https://golang.org/doc/install/source#environment.


