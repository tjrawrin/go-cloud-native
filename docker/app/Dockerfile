# Build environment
# -----------------
FROM golang:1.14-alpine as build-env
WORKDIR /go-cloud-native

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate


# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash mysql-client

COPY --from=build-env /go-cloud-native/bin/app /go-cloud-native/
COPY --from=build-env /go-cloud-native/bin/migrate /go-cloud-native/
COPY --from=build-env /go-cloud-native/migrations /go-cloud-native/migrations

COPY --from=build-env /go-cloud-native/docker/app/bin /usr/local/bin/go-cloud-native/
RUN chmod +x /usr/local/bin/go-cloud-native/*

EXPOSE 8080
CMD ["/go-cloud-native/app"]
