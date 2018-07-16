.PHONY: cli

cli:
	go build -o stickergen cmd/stickergen/main.go

js:
	gopherjs build --minify
