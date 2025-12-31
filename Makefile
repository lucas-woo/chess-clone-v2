lint-breaking:
	go tool buf breaking --against 'https://github.com/lucas-woo/chess-clone-v2'

lint-proto:
	go tool buf lint --config buf.yaml

generate-proto:
	go tool buf generate --template buf.gen.yaml