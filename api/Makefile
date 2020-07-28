install:
	go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	rm -rf ./lib
	mkdir lib
	curl -sSL https://git.io/get-mo -o mo
	mv mo lib/
	npm install

generate:
	./scripts/openapi-release.sh