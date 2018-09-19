
buildDir = ./build
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

build: clean build-linux build-windows build-darwin build-arm

build-darwin:
	env \
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
		GOOS=windows \
		GOARCH=amd64 \
	go build \
		-v \
		-p 8 \
		-o "$(buildDir)/$(projectName)-$(version)-Windows-x86_64" \
		-ldflags "-X main.GitRev=$(gitRev)" \

build-arm:
	env \
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
	mkdir -p "$(buildDir)/log"

	go test \
		-coverprofile="$(buildDir)/log/coverage.out" \
		./...

	go tool cover \
		-html="$(buildDir)/log/coverage.out" \
		-o "$(buildDir)/log/coverage.html"
