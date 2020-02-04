
compile:
	go build
	./limiter

test:
	go test -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html
	@echo "Point your browser at `pwd`/coverage.html"
