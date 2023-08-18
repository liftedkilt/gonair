TARGET = gonair
GO = go
LIPO = lipo

# Directories for intermediate files
BUILD_DIR = build
ARM64_DIR = $(BUILD_DIR)/arm64
AMD64_DIR = $(BUILD_DIR)/amd64

.PHONY: all clean

all: $(TARGET)

$(TARGET): $(ARM64_DIR)/$(TARGET) $(AMD64_DIR)/$(TARGET)
	$(LIPO) -create -output $@ $^

$(ARM64_DIR)/$(TARGET): 
	mkdir -p $(ARM64_DIR)
	GOOS=darwin GOARCH=arm64 $(GO) build -o $@ .

$(AMD64_DIR)/$(TARGET): 
	mkdir -p $(AMD64_DIR)
	GOOS=darwin GOARCH=amd64 $(GO) build -o $@ .

clean:
	rm -rf $(BUILD_DIR) $(TARGET)

