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

run-profile-server: cmd/servers/puzzle/main.go
	go run cmd/servers/profile/main.go

run-client: cmd/client/main.go
	clear;
	GIN_MODE=release go run cmd/client/main.go;

run-auth-test: cmd/servers/test/auth/main.go
	go run cmd/servers/test/auth/main.go;

run-admin-test: cmd/servers/test/admin/main.go
	go run cmd/servers/test/admin/main.go;

run-profile-test: cmd/servers/test/admin/main.go
	go run cmd/servers/test/profile/main.go;