

.PHONY: setup
setup:
	go get -u github.com/golang/dep/cmd/dep
	$(MAKE) dep

.PHONY: dep
dep:
	@dep ensure -v

.PHONY: clean
clean:
	rm -f deputy

.PHONY: build
build: clean ## build the executable
	go build

.PHONY: start
start: build
	./deputy

.PHONY: test
test:
	go test