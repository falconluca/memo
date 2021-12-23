.PHONY: protoc
protoc:
	@if ! which protoc > /dev/null; then \
		echo "protoc not install" >&2; \
		exit 1; \
	fi
	@if ! which protoc-gen-go > /dev/null; then \
  		echo "protoc-gen-go not install" >&2; \
		exit 1; \
	fi
	@if ! which protoc-gen-go-grpc > /dev/null; then \
  		echo "protoc-gen-go-grpc not install" >&2; \
	fi
	for file in $$(git ls-files '*.proto'); do \
  		protoc -I $$(dirname $$file) \
  		--go_out=:$$(dirname $$file) --go_opt=paths=source_relative \
  		--go-grpc_out=:$$(dirname $$file) --go-grpc_opt=paths=source_relative \
  		$$file; \
	done

.PHONY: wire
wire:
	wire ./cmd/server