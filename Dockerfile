# Update Dockerfile
FROM gcr.io/distroless/static
ENTRYPOINT ["/starlink_exporter"]
COPY starlink_exporter /
