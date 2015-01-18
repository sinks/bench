

install: build
	go install

build: version
	go build

version: get sqlite

sqlite:
	git -C $GOAPTH/src/github.com/mattn/go-sqlite3/ checkout c9a0db5d8951646743317f0756da0339fe144dd5

get:
	go get -u
