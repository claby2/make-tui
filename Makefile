PROJECT = make-tui
BUILD_DIR ?= build
APP_SOURCES = parser.go \
			  render.go \
			  main.go \

build: $(APP_SOURCES)
	go build -o $(BUILD_DIR)/maketui $(APP_SOURCES)
.PHONY: build

run: $(APP_SOURCES)
	go run $(APP_SOURCES)
.PHONY: run

test:
	go test ./...
.PHONY: test

clean:
	rm -rf $(BUILD_DIR)
.PHONY: clean
