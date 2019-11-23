GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=${GOCMD} run
GOMODULE=${GOCMD} mod

BUILD_TOOLS=.build/tools
PACKER_URL=https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz
PACKER=.build/tools/upx

VERSION?=SNAPSHOT_$(shell git describe --always --long HEAD)

LINKER_VARS=-ldflags "-s -w -X main.Version=${VERSION}"

GOBUILD=$(GOCMD) build ${LINKER_VARS}

LINUX_AMD64_BIN=${DIST_PATH}/${BINARY_NAME}-linux-amd64
WIN_AMD64_BIN=${DIST_PATH}/${BINARY_NAME}-win-amd64.exe
DARWIN_AMD64_BIN=${DIST_PATH}/${BINARY_NAME}-darwin-amd64

DIST_PATH=dist
BINARY_NAME=yoke
YOKE_MAIN_SOURCE_PATH=cmd/yoke/main.go

# default task
all: test build

# build for current os and arch
build: 
	$(GOBUILD) -o ${DIST_PATH}/yoke ${YOKE_MAIN_SOURCE_PATH}

# run tests
test: 
	$(GOTEST) -v ./...

# cleans the directories
clean: 
	$(GOCLEAN) ${YOKE_MAIN_SOURCE_PATH}
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -rf ${DIST_PATH}
	rm -rf ${BUILD_TOOLS}

# downloads required tools
dl-tools:
	mkdir -p ${BUILD_TOOLS}
	curl -Lso ${BUILD_TOOLS}/upx.tar.xz ${PACKER_URL}
	echo "b5d6856b89dd696138ad8c7245a8f0dae4b76f41b5a31c7c43a21bc72c479c4e  ${BUILD_TOOLS}/upx.tar.xz" | sha256sum --check
	tar -C ${BUILD_TOOLS} --strip=1 -xf ${BUILD_TOOLS}/upx.tar.xz

# cross compile sources
crosscompile:
	GOOS=linux GOARCH=amd64 ${GOBUILD} -o ${LINUX_AMD64_BIN} ${YOKE_MAIN_SOURCE_PATH}
	GOOS=windows GOARCH=amd64 ${GOBUILD} -o ${WIN_AMD64_BIN} ${YOKE_MAIN_SOURCE_PATH}
	GOOS=darwin GOARCH=amd64 ${GOBUILD} -o ${DARWIN_AMD64_BIN} ${YOKE_MAIN_SOURCE_PATH}

# pack executables
pack:
	${PACKER} ${LINUX_AMD64_BIN} ${WIN_AMD64_BIN} ${DARWIN_AMD64_BIN}

dist: crosscompile pack
