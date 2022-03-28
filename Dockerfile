FROM golang:1.18 as build
WORKDIR $GOPATH/src/github.com/ahmetozer/distroless-helper/
COPY ["main.go", "go.mod", "go.sum", "./"]

RUN go mod tidy
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /bin/distroless-helper
RUN apt update; apt install --no-install-recommends -y upx
RUN upx --brute /bin/distroless-helper

FROM scratch
COPY --from=build /bin/distroless-helper /bin/distroless-helper
LABEL org.opencontainers.image.source="https://github.com/ahmetozer/distroless-helper"
ENTRYPOINT ["/bin/distroless-helper"]