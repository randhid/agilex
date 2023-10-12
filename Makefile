
limobase: *.go 
	go build -o limobase *.go

test:
	go test

lint:
	gofmt -w -s .

module: limobase 
	tar czf module.tgz limobase
