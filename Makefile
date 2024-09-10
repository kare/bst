
.PHONY: clean
clean:
	@rm -f prof bst.test bench.txt

.PHONY: benchmark
benchmark: clean
	go test -c
	./bst.test -test.bench=. -test.count=5 >bench.txt
	@cat bench.txt

.PHONY: mem-profile
mem-profile: clean
	go test -run=XXX -bench=. -memprofile=prof kkn.fi/bst
	go tool pprof bst.test prof

.PHONY: cpu-profile
cpu-profile: clean
	go test -run=XXX -bench=. -cpuprofile=prof kkn.fi/bst
	go tool pprof bst.test prof

.PHONY: build
build:
	go build kkn.fi/bst

.PHONY: test
test:
	go test kkn.fi/bst
