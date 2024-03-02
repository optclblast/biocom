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

# Generate events protos.
.SILENT:
.PHONY: events.proto
events.proto:
	find pkg/proto/events -iname *.proto -exec protoc -I=pkg/proto \
		--go_out=paths=source_relative:./pkg/proto/gen \
		--go-grpc_out=paths=source_relative:./pkg/proto/gen {} \;

# Generate Garganta API protos.
.SILENT:
.PHONY: garganta.proto
garganta.proto:
	find pkg/proto/garganta -iname *.proto -exec protoc -I=pkg/proto \
		--go_out=paths=source_relative:./pkg/proto/gen \
		--go-grpc_out=paths=source_relative:./pkg/proto/gen {} \;

# Generate Cernunnos API protos.
.SILENT:
.PHONY: cernunnos.proto
cernunnos.proto:
# Go code
	find pkg/proto/cernunnos -iname *.proto -exec protoc -I=pkg/proto \
		--go_out=paths=source_relative:./pkg/proto/gen \
		--go-grpc_out=paths=source_relative:./pkg/proto/gen {} \;
# Elixir code
	find pkg/proto/cernunnos -iname *.proto -exec protoc -I=pkg/proto \
		--elixir_out=plugins=grpc:./internal/services/cernunnos/lib {} \;