lint-breaking:
	go tool buf breaking --against 'https://github.com/lucas-woo/chess-clone-v2'

lint-proto:
	go tool buf lint --config buf.yaml

generate-proto:
	go tool buf generate --template buf.gen.yaml

run-puzzle-server: cmd/servers/puzzle/main.go
	go run cmd/servers/puzzle/main.go

run-auth-server: cmd/servers/auth/main.go;
	go run cmd/servers/auth/main.go

run-client: cmd/client/main.go
	clear;
	GIN_MODE=release go run cmd/client/main.go;

run-auth-test: cmd/servers/test/auth/main.go
	go run cmd/servers/test/auth/main.go;

run-puzzle-test: cmd/servers/test/puzzle/main.go
	go run cmd/servers/test/puzzle/main.go;