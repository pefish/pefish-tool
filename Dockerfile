FROM pefish/ubuntu-go:v1.22 AS builder
WORKDIR /app
ENV GO111MODULE=on
COPY ./ ./
RUN make

FROM pefish/ubuntu:v20.04
WORKDIR /app
COPY --from=builder /app/build/bin/linux/ /app/bin/
ENV GO_CONFIG=/app/config/config.yaml
CMD ["/app/bin/pefish-tool"]

# docker build --progress=plain -t pefish/pefish-tool:v0.0.1 .
# docker run -ti --name pefish-tool pefish/pefish-tool:v0.0.1
