.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/wire/cmd/wire@latest

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
  		exit 1; \
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

.PHONY: docker-build
docker-build:
	docker build -t shaohsiung/memo:latest -f ./Dockerfile .

.PHONY: docker-push
docker-push:
	docker push shaohsiung/memo:latest

.PHONY: kube-deploy-mysql
kube-deploy-mysql:
	kubectl apply -f k8s/mysql.yaml
	kubectl rollout status deploy/mysql

.PHONY: kube-deploy-memo
kube-deploy-mysql:
	kubectl apply -f k8s/memo.yaml
	kubectl rollout status deploy/memo

.PHONY: kube-port-forward
kube-port-forward:
	kubectl port-forward $$(kubectl get pod -l app=memo -o jsonpath="{.items[0].metadata.name}") 50051:50051

.PHONY: kube-deploy-all
kube-deploy-all:
	make kube-deploy-mysql
	make kube-deploy-memo
	make kube-port-forward

.PHONY: kube-delete-all
kube-delete-all:
	kubectl delete -f k8s/