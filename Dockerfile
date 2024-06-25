FROM golang:1.22 as build
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/server .
CMD ["./server"]