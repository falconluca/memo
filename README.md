# Memo microservice

**Makefile**

| 命令 | 说明 |
| ---- | ---- |
| `make init` | 安装 protoc-gen-go, protoc-gen-go-grpc 和 wire |
| `make protoc` | 根据 proto 文件生成 *.pb.go 和 *_grpc.pb.go |
| `make wire` | 依赖注入, 根据 wire.go 文件生成 wire_gen.go |
| `make docker-build` | 构建 docker 镜像 |
| `make kube-deploy-mysql` | 在 k8s 集群中部署 MySQL 服务 |
| `make kube-deploy-memo` | 在 k8s 集群中部署 Memo 服务 |
| `make kube-port-forward` | 开启 Memo 服务测试入口 |
| `make kube-deploy-all` | 在 k8s 集群中部署所有资源 |
| `make kube-delete-all` | 在 k8s 集群中删除所有资源 |


**使用 grpcurl 访问 grpc 服务**

```shell
grpcurl -plaintext 127.0.0.1:50051 list 
# Output:
# api.protobuf.Memo
# grpc.health.v1.Health
# grpc.reflection.v1alpha.ServerReflection

grpcurl -d '{"item": {"title": "记得吃早餐", "description": "咖啡配奶酪", "remind_at": "2021-12-23T10:04:05.999999999Z"}}' \
  -plaintext 127.0.0.1:50051 api.protobuf.Memo.Create
# Output:
# {
#   "id": "1"
# }

grpcurl -d '{"id": 1}' -plaintext 127.0.0.1:50051 api.protobuf.Memo.Get
# Output:
# {
#   "item": {
#     "id": "1",
#     "title": "记得吃早餐",
#     "description": "咖啡配奶酪",
#     "remindAt": "2021-12-23T10:04:06Z"
#   }
# }

grpcurl -plaintext 127.0.0.1:50051 api.protobuf.Memo.List
# Output:
# {
#   "items": [
#     {
#       "id": "1",
#       "title": "记得吃早餐",
#       "description": "咖啡配奶酪",
#       "remindAt": "2021-12-23T10:04:06Z"
#     }
#   ]
# }

grpcurl -d '{"item": {"title": "记得吃早餐", "description": "牛奶配奶酪", "remind_at": "2021-12-23T10:04:05.999999999Z"}}' \
  -plaintext 127.0.0.1:50051 api.protobuf.Memo.Update
# Output:
# {
#   "updated": true
# }

grpcurl -d '{"id": 1}' -plaintext 127.0.0.1:50051 api.protobuf.Memo.Delete
# Output: 
# {
#   "deleted": true
# }
```