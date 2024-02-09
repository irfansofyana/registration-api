.PHONY: clean all init generate generate_mocks keys

all: build/main

build/main: cmd/main.go generated
	@echo "Building..."
	go build -o $@ $<

clean:
	rm -rf generated private.pem public.pem

init: generate keys
	go mod tidy
	go mod vendor

test:
	go test -short -coverprofile coverage.out -v ./...

generate: generated generate_mocks

keys:
	openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -pubout -in private.pem -out public.pem

generated: api.yml
	@echo "Generating files..."
	mkdir generated || true
	oapi-codegen --package generated -generate types,server,spec $< > generated/api.gen.go

INTERFACES_GO_FILES := $(shell find repository -name "interfaces.go")
INTERFACES_GEN_GO_FILES := $(INTERFACES_GO_FILES:%.go=%.mock.gen.go)

generate_mocks: $(INTERFACES_GEN_GO_FILES)
$(INTERFACES_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))