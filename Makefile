BINARY = dolister
COVERAGE_OUT = coverage.out
COVERAGE_HTML = coverage.html

GO = go

all: build

build:
	$(GO) build -o $(BINARY) cmd/main.go

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)
	rm -f $(COVERAGE_OUT)
	rm -f $(COVERAGE_HTML)

test:
	$(GO) test -v -cover ./...

coverage:
	$(GO) test -coverprofile=$(COVERAGE_OUT) ./...
	$(GO) tool cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)

.PHONY: all run clean test coverage
