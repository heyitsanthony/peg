peg.peg.go: peg.peg
	peg peg.peg

peg-canon: peg.peg.go main.go
	go build -i -v

test: peg-canon
	cat peg.peg | peg-canon >1.peg
	cat ../../peg.peg | peg-canon >2.peg
	diff 1.peg 2.peg