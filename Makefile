build: deps
	go build calculator/server/server.go

clean:
	go clean -i google.golang.org/grpc/...

deps:
	dep ensure

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate google.golang.org/grpc/...

test: testdeps
	go test -cpu 1,4 -timeout 7m google.golang.org/grpc/...

.PHONY: \
	all \
	build \
	clean \
	deps \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testdeps \
	testrace \
	updatedeps \
	updatetestdeps \
	vet \
	vetdeps