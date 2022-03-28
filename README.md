# Distroless Helper

Copy binary and libraries to the target folder.

Example usage

```Dockerfile
FROM debian as base
COPY --from=ghcr.io/ahmetozer/distroless-helper /bin/distroless-helper /bin/distroless-helper
RUN distroless-helper /usr/sbin/nginx /opt

RUN cp -a --parents /etc/nginx /opt/ && \
 cp -a --parents /var/lib/nginx/logs/ /opt/ && \
 cp -a --parents /etc/passwd /opt/ && \
 cp -a --parents /etc/group /opt/ && \
 cp -a --parents /var/lib/nginx/tmp/ /opt/ && \
 cp -a --parents /var/log/nginx/ /opt/ &&\
 cp -a --parents /run/nginx/ /opt/

FROM scratch
COPY --from=base /opt/ /
ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
```

For more example, you can visit [ahmetozer/containers](github.com/ahmetozer/containers)
