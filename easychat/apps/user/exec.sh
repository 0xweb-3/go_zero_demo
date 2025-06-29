# user rpc 模版代码生成，在easychat目录下
goctl rpc protoc ./apps/user/rpc/user.proto --go_out=./apps/user/rpc/ --go-grpc_out=./apps/user/rpc/ --zrpc_out=./apps/user/rpc/

# 生成数据模型，在easychat下执行
goctl model mysql ddl -src="./deploy/sql/user.sql" -dir="./apps/user/models/" -c

# 构建api服务
goctl api go -api apps/user/api/user.api -dir apps/user/api -style gozero





