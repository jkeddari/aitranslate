templ:
	templ generate

tailwindcss:
	tailwindcss -i input.css -o static/css/styles.min.css --minify

dev:
	air .

#build
build: tailwindcss templ
	go build -o build/ ./...

docker:
	docker build -t aitranslate .
