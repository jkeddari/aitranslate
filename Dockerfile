# Builder stage
FROM golang:latest AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . /app
RUN templ generate && CGO_ENABLED=0 go build -o bin/ ./...

# Release stage
FROM gcr.io/distroless/static-debian11
COPY --from=build /app/bin/server /server
CMD ["/server"]
