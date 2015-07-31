TYPE=patch

build::
	mkdir -p build
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 godep go build -o build/emp
	tar -C build -czf build/emp-linux-amd64-`cat VERSION`.tgz emp
	rm build/emp
	CGO_ENABLED=0 godep go build -o build/emp
	tar -C build -czf build/emp-mac-amd64-`cat VERSION`.tgz emp
	rm build/emp

bump:
	bumpversion ${TYPE}
