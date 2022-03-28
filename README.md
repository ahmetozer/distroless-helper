# Distroless Helper

Copy binary and libraries to the target folder.

Example usage

```Dockerfile
FROM debian
COPY --from=ghcr.io/ahmetozer/distroless-helper /bin/distroless-helper /bin/distroless-helper
RUN distroless-helper /usr/sbin/nginx /target

RUN cp -a --parents /etc/nginx /target/ && \
 cp -a --parents /var/lib/nginx/logs/ /target/ && \
 cp -a --parents /etc/passwd /target/ && \
 cp -a --parents /etc/group /target/ && \
 cp -a --parents /var/lib/nginx/tmp/ /target/ && \
 cp -a --parents /var/log/nginx/ /target/ &&\
 cp -a --parents /run/nginx/ /target/

FROM scratch
COPY --from=builder /target/ /
ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
```

For more example, you can visit [ahmetozer/containers](github.com/ahmetozer/containers)
