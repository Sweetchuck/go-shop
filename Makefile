
buildDir = ./build
logDir = ./log
projectName = shop
versionBase = 1.0.0
gitRev = "$$(git rev-parse --short=7 HEAD)"
version = "${versionBase}-${gitRev}"

.PHONY: clean build build-darwin build-linux build-windows

clean:
	[ ! -f $(projectName) ] || rm $(projectName);

	if [ -d $(buildDir) ]; then \
		find $(buildDir) -mindepth 1 -maxdepth 1 -exec rm -rf {} \; ; \
	fi;

	if [ -d $(logDir) ]; then \
		find $(logDir) -mindepth 1 -maxdepth 1 -exec rm -rf {} \; ; \
	fi;

build: clean build-linux build-windows build-darwin build-arm

build-darwin:
	env \
		CGO_ENABLED=1 \
		GOOS=darwin \
		GOARCH=amd64 \
	go build \
		-v \
		-p 8 \
		-o "$(buildDir)/$(projectName)-$(version)-Darwin-x86_64" \
		-ldflags "-X main.GitRev=$(gitRev)" \
		.

build-linux:
	env \
		CGO_ENABLED=0 \
		GOOS=linux \
		GOARCH=amd64 \
	go build \
		-v \
		-p 8 \
		-o "$(buildDir)/$(projectName)-$(version)-Linux-x86_64" \
		-ldflags "-X main.GitRev=$(gitRev)" \
		.

build-windows:
	env \
		CGO_ENABLED=1 \
		GOOS=windows \
		GOARCH=amd64 \
	go build \
		-v \
		-p 8 \
		-o "$(buildDir)/$(projectName)-$(version)-Windows-x86_64" \
		-ldflags "-X main.GitRev=$(gitRev)" \

build-arm:
	#		CXX_FOR_TARGET=CC_FOR_linux_arm \
	env \
		CC=arm-suse-linux-gnueabi-gcc \
		CC_FOR_TARGET=CC_FOR_linux_arm \
		CC_FOR_linux_arm=1 \
		CGO_ENABLED=1 \
		CGO_CFLAGS="--sysroot='/usr/lib64/gcc/arm-suse-linux-gnueabi/8'" \
		CGO_LDFLAGS=""\
		GOOS=linux \
		GOARCH=arm \
		GOARM=6 \
	go build \
		-v \
		-p 8 \
		-o "$(buildDir)/$(projectName)-$(version)-armv6l" \
		-ldflags "-X main.GitRev=$(gitRev)" \
		.

test:
	mkdir -p "$(logDir)"

	go test \
		-coverprofile="$(logDir)/coverage.out" \
		./...

	go tool cover \
		-html="$(logDir)/coverage.out" \
		-o "$(logDir)/coverage.html"
