build:
	GOOS=linux GOARCH=amd64 go build -o _build/linux github.com/jjanvier/tdd

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o _build/macos github.com/jjanvier/tdd

install:
	sudo cp /usr/local/bin/tdd /usr/local/bin/tdd.bak
	sudo mv ./_build/linux/tdd /usr/local/bin/tdd

release:
	./bin/prepare_github_package.sh
