.PHONY: all
TARGET := gateway
GOENV := GOOS=linux GOARCH=amd64
all:
	CGO_ENABLED=0 ${GOENV} go build -o ./bin/${TARGET}

clean:
	rm -rf ${TARGET}

format:
	gofmt -w ./
	goimports -w ./
	goint ./...
build:
	CGO_ENABLED=0 ${GOENV} go build -buildvcs=false -o ./bin/${TARGET}