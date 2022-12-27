FROM golang:alpine AS build

RUN apk add --update git
# GOPROXY resolves dependencies treefrom cache or repository
ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src/github.com/ArthurQR98/go-hexagonal_http_api_challenge
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/challenge-fiber-api cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/challenge-fiber-api /go/bin/challenge-fiber-api
ENTRYPOINT ["/go/bin/challenge-fiber-api"]
