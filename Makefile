TARGET=random_user
TARGET_BUILD_REVISION=`git rev-parse --short HEAD`
LDFLAGS="-X main.buildRevision=$(TARGET_BUILD_REVISION) -s -w"

all: build
build:
	go build -ldflags=$(LDFLAGS) -o $(TARGET) main.go;

.PHONY: run
run:
	go run $(TARGET)

.PHONY: test
test:
	go test -race -v

.PHONY: clean
clean:
	rm -f $(TARGET)

.PHONY: style-fix
style-fix:
	gofmt -w .

.PHONE: lint
lint:
	golangci-lint run

.PHONY: download
download:
	go mod download

.PHONY: upgrade
upgrade:
	go mod tidy
	go get -u all ./...
