services.proto:
	make ws.proto && make warden.proto && make events.proto

# Generate WS API protos
.SILENT:
.PHONY: ws.proto
ws.proto:
	find pkg/proto -iname *.proto -exec protoc -I=pkg/proto \
		--go_out=paths=source_relative:pkg/proto/gen \
		--go-grpc_out=paths=source_relative:pkg/proto/gen {} \;


# Generate warden API protos.
.SILENT:
.PHONY: warden.proto
warden.proto:
	find pkg/proto/warden -iname *.proto -exec protoc -I=pkg/proto \
		--go_out=paths=source_relative:./pkg/proto/gen \
		--go-grpc_out=paths=source_relative:./pkg/proto/gen {} \;

events.proto:
	find pkg/proto/events -iname *.proto -exec protoc -I=pkg/proto \
		--go_out=paths=source_relative:./pkg/proto/gen \
		--go-grpc_out=paths=source_relative:./pkg/proto/gen {} \;

garganta.proto:
	find pkg/proto/garganta -iname *.proto -exec protoc -I=pkg/proto \
		--go_out=paths=source_relative:./pkg/proto/gen \
		--go-grpc_out=paths=source_relative:./pkg/proto/gen {} \;