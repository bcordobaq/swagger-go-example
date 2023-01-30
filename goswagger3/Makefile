.PHONY: swagger

run:
	go run .

install_deps:
	go install github.com/parvez3019/go-swagger3@latest

api-spec:
	go-swagger3 --debug --module-path . --main-file-path ./cmd/main.go --output oas.json --schema-without-pkg --generate-yaml true
