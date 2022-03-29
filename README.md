# Distroless Helper

Copy binary and libraries to the target folder.

Example usage

```Dockerfile
FROM debian as base
COPY --from=ghcr.io/ahmetozer/distroless-helper /bin/distroless-helper /bin/distroless-helper
RUN /bin/distroless-helper /bin/bash /opt

FROM scratch
COPY --from=base /opt/ /
USER 65534
ENTRYPOINT ["/bin/bash"]
```

For more example, you can visit [ahmetozer/containers](https://github.com/ahmetozer/containers)
