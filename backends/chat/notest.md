

 protoc /api/v1/*.proto \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    --proto_path=.

grpcurl -plaintext -d  '{ "message": "christmas eve bike class", "time" : "time" }' localhost:8080 api.v1.Chat/Message