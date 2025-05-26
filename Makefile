
dev:
	make -j3 tailwind templ server

tailwind:
	./tailwindcss --cwd ./web/public/css -i input.css -o styles.css -m --watch

templ:
	go tool templ generate --watch

server:
	air

sqlc:
	go tool sqlc generate
