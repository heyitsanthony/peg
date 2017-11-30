golang: main.go go.peg.go
	go build -i -v
go.peg.go: go.peg
	peg go.peg

clean:
	rm -f go.peg.go golang

.PHONY: bench
bench: golang
	./bench.sh /usr/lib/go/src/*/{*,*/*,*/*/*}.go >bench.out

