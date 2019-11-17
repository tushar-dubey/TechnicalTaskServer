
build: deps
	go build ./...

clean:
	go clean -i ./...

deps:
	go get -d -v ./...


proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate ./...


test: testdeps
	go test ./...

testdeps:
	go get -d -v -t ./...