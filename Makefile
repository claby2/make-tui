PROJECT = make-tui
BUILD_DIR ?= build
APP_SOURCES = cmd/parser.go \
			  cmd/render.go \
			  cmd/main.go \

build: $(APP_SOURCES)
	go build -o $(BUILD_DIR)/maketui $(APP_SOURCES)
.PHONY: build

run: $(APP_SOURCES)
	go run $(APP_SOURCES)
.PHONY: run

clean:
	rm -rf $(BUILD_DIR)
.PHONY: clean
