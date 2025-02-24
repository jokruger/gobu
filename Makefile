test:
	go test ./tests/unit

bench:
	go test -bench=. ./tests/benchmark
