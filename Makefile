BINARY := dt
INSTALL_DIR := $(HOME)/bin

.PHONY: all build install clean

all: build

build:
	@echo "Building $(BINARY)"
	@go build -o $(BINARY)

install: build
	@echo "Installing $(BINARY) to $(INSTALL_DIR)"
	@mkdir -p $(INSTALL_DIR)
	@install -Dm755 $(BINARY) $(INSTALL_DIR)/$(BINARY)

clean:
	@echo "Cleaning $(BINARY)"
	@rm -f $(BINARY)