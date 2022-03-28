FROM golang:1.18 as build
WORKDIR $GOPATH/src/github.com/ahmetozer/distroless-helper/
COPY ["main.go", "go.mod", "go.sum", "./"]

RUN go mod tidy
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /bin/distroless-helper

FROM ghcr.io/ahmetozer/containers/upx as upx
COPY --from=build /bin/distroless-helper /bin/distroless-helper
RUN upx --brute /bin/distroless-helper

FROM scratch
COPY --from=upx /bin/distroless-helper /bin/distroless-helper
LABEL org.opencontainers.image.source="https://github.com/ahmetozer/distroless-helper"
ENTRYPOINT ["/bin/distroless-helper"]