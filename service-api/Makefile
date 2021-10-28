generate:
	mkdir -p OpenAPI

	rm -r proto/pb | true
	mkdir -p proto/pb
	mkdir -p proto/pb/pb_models
	mkdir -p proto/pb/pb_auth
	mkdir -p proto/pb/pb_post
	mkdir -p proto/pb/pb_comment
	mkdir -p proto/pb/pb_notify

	protoc -I proto \
	--go_out=plugins=grpc,paths=source_relative:./proto/pb/pb_models \
	--grpc-gateway_out=paths=source_relative:./proto/pb/pb_models \
	--openapiv2_out=OpenAPI/ proto/models.proto

	protoc -I proto \
	--go_out=plugins=grpc,paths=source_relative:./proto/pb/pb_auth \
	--grpc-gateway_out=paths=source_relative:./proto/pb/pb_auth \
	--openapiv2_out=OpenAPI/ proto/auth_service.proto

	protoc -I proto \
	--go_out=plugins=grpc,paths=source_relative:./proto/pb/pb_post \
	--grpc-gateway_out=paths=source_relative:./proto/pb/pb_post \
	--openapiv2_out=OpenAPI/ proto/post_service.proto

	protoc -I proto \
	--go_out=plugins=grpc,paths=source_relative:./proto/pb/pb_comment \
	--grpc-gateway_out=paths=source_relative:./proto/pb/pb_comment \
	--openapiv2_out=OpenAPI/ proto/comment_service.proto

		protoc -I proto \
	--go_out=plugins=grpc,paths=source_relative:./proto/pb/pb_notify \
	--grpc-gateway_out=paths=source_relative:./proto/pb/pb_notify \
	--openapiv2_out=OpenAPI/ proto/notify_service.proto

	# Generate static assets for OpenAPI UI
	statik -m -f -src OpenAPI/


goInstall:
	go install

install:
	go get github.com/golang/protobuf/protoc-gen-go 
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway 
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 
	go get github.com/rakyll/statik 
	go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators 
	go get github.com/stretchr/testify 
	go get github.com/jinzhu/copier 
	go get go.mongodb.org/mongo-driver/mongo
	go get github.com/joho/godotenv
	go get github.com/tidwall/gjson
	go get github.com/streadway/amqp
	go get github.com/go-redis/redis/v8
	go get github.com/dgrijalva/jwt-go

clean:
	rm -r proto/pb

run:
	go run main.go

test:
	go test -cover -race ./...

gencert:
	cd cert; ./gen.sh; cd ..