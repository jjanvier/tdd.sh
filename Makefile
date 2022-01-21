build:
	go build -o _build github.com/jjanvier/tdd

install:
	sudo cp /usr/local/bin/tdd /usr/local/bin/tdd.bak
	sudo mv ./_build/tdd /usr/local/bin/tdd

release:
	./bin/prepare_github_package.sh
