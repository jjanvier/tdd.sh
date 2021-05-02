build:
	go build

install:
	sudo cp /usr/local/bin/tdd /usr/local/bin/tdd.bak
	sudo mv tdd /usr/local/bin/tdd

release:
	./bin/prepare_github_package.sh
