GO_BUILDER_VERSION=latest
PRIVATE_KEY=$(shell cat ~/.gnupg/pubring.gpg | base64)

release-local:
	goreleaser release -f .goreleaser.local.yaml --rm-dist

build-local:
	goreleaser build -f .goreleaser.local.yaml --snapshot --rm-dist

build-gythialy:
	docker run --rm --privileged \
		-e PRIVATE_KEY="$(PRIVATE_KEY)" \
		-v $(CURDIR):/golang-cross-example \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v $(GOPATH)/src:/go/src \
		-w /golang-cross-example \
		ghcr.io/gythialy/golang-cross:$(GO_BUILDER_VERSION) --snapshot --rm-dist

