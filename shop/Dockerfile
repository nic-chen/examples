ARG BUILDPLATFORM=amd64

FROM --platform=$BUILDPLATFORM golang:1.19 as api-builder

ARG ENABLE_PROXY=false
ARG TARGETOS
ARG TARGETARCH

WORKDIR /usr/local/shop
COPY ./ ./

RUN if [ "$ENABLE_PROXY" = "true" ] ; then go env -w GOPROXY=https://goproxy.io,direct ; fi \
    && go env -w GO111MODULE=on \
    && CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o ./shop .

FROM debian:bullseye-slim as prod

WORKDIR /usr/local/shop

COPY --from=api-builder /usr/local/shop/shop ./shop

CMD ["/usr/local/shop/shop"]
