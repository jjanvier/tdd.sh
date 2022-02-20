# Releasing

[goreleaser](https://goreleaser.com) is used to under the hood build packages and to deploy them on Github via a Release.

## Release process

1. launch `release-local`:
   - this command removes the `_build` directory before proceeding to have a clean build directory
   - it also builds the packages that can be built locally (all Windows versions and Linux amd64)
   - finally, it uploads those package on Github and creates a new release
2. launch `make build-gythialy`:
   - this command builds the Darwin packages
3. upload manually the packages that have been built on step 2

## Explanations

The problem is that releasing for various platforms and architectures is crazy difficult...

### About CGO

To play the notification bell sound, _TDD.sh_ has a dependency on [hajimehoshi/oto](https://github.com/hajimehoshi/oto). This library [requires](https://github.com/hajimehoshi/oto#prerequisite) `libasound2-dev` on Linux for instance, so that it can call some C code: [alsa/asoundlib.h](https://github.com/hajimehoshi/oto/blob/main/driver_unix.go#L22) to be exact. Same for other platforms that may have other requirements.

That means that when we package _TDD.sh_ for Linux, we must embed somehow `alsa/asoundlib.h`. And here comes the [CGO](https://go.dev/blog/cgo). It's the ability to embed native C code, by compiling it, during the packaging of a Go build. So yes, that means we'll use `gcc` during the build of our Go package.

For _goreleaser_ to enable CGO during compilation, we must use `CGO_ENABLED=1` as environment variable. 

### Building locally

My current computer is a standard amd64 Linux with `libasound2-dev` installed. Therefore, I'm able to build a _TDD.sh_ package for `linux-amd64`. And that's it...

I could compile it for Windows if I'd install some [others gcc compilers](https://dh1tw.de/2019/12/cross-compiling-golang-cgo-projects/#get-a-cross-compiler) like `gcc-mingw-w64-x86-64` or `gcc-multilib` for instance. I could also build it for _arm64_ if I'd install other compilers and `libasound2-dev:arm64` on my laptop. But of course, I don't want to do that.

An example of such a build would be:
```
GOOS=windows GOARCH=386 CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc go build
```
where `CC` is the `gcc` compiler and `CXX` the `g++` compiler.

### Using an existing Dockerized toolchain?

To compile for other platforms and architectures, I found 2 Dockerized toolchains so far:
- https://github.com/gythialy/golang-cross
- https://github.com/goreleaser/goreleaser-cross

`gythialy/golang-cross` is the one I've been trying the most. It seems to work fine. My problem is that so far I'm not able to customize the Docker images to install `libasound2-dev`. See this [Github ticket](https://github.com/gythialy/golang-cross/issues/88) and `~/code/oss/golang-cross`.

To know which compiler to use on which architecture/platform, check [this](https://github.com/goreleaser/goreleaser-cross#supported-toolchainsplatforms).

### Using gccgo?

Maybe an alternative to look at is [gccgo](https://go.dev/doc/install/gccgo). 

## Enhancements

TODO
- use only one Dockerized toolchain to build all the packages would be the best
- that would also allow using only one `.goreleaser.yaml` config file
