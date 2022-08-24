run-local:
	go run ./cmd/cli/*.go

run-docker:
	docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.18.5-alpine3.16 go run ./cmd/cli/*.go