# Build environment
# -----------------
FROM golang:1.14-alpine as build-env
WORKDIR /go-cloud-native
RUN apk update && apk add --no-cache gcc musl-dev git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash
COPY --from=build-env /go-cloud-native/bin/app /go-cloud-native/
EXPOSE 8080
CMD ["/go-cloud-native/app"]
