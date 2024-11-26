# syntax = docker/dockerfile:1

ARG ALPINE_VERSION=3.20
ARG GOLANG_VERSION=1.23.3
ARG GINKGO_VERSION=v2.22.0

# ====== Base stage ======
FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS base

ARG GINKGO_VERSION

RUN apk --no-cache add \
    make \
    git \
    gcc \
    musl-dev \
    ca-certificates

RUN go install github.com/onsi/ginkgo/v2/ginkgo@${GINKGO_VERSION}

WORKDIR /go/src/github.com/NishimuraTakuya-nt/go-rest-chi


# ====== Test stage ======
FROM base AS test

COPY go.mod go.sum Makefile ./
RUN --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    make go-download

COPY . .
RUN --mount=type=cache,id=go-test,target=/root/.cache/go-build \
    --mount=type=bind,target=/output \
    mkdir -p ./coverage && \
    make test-ginkgo-coverage
# id このままで良いのか？ modでいいのか？
RUN touch aaa.txt
RUN ls -la
RUN ls -la coverage
RUN cat coverage/coverage.txt
RUN mkdir -p /output
RUN cp coverage/coverage.txt /output/
RUN ls -la /output
RUN ls -la /

# ====== Build stage ======
FROM base AS builder

COPY go.mod go.sum Makefile ./
RUN --mount=type=cache,id=go-rest-chi-pkg,target=/go/pkg \
    --mount=type=secret,id=netrc,target=/root/.netrc \
    make go-download

COPY . .
RUN --mount=type=cache,id=go-rest-chi-pkg,target=/go/pkg \
    --mount=type=cache,id=go-rest-chi-go-build,target=/root/.cache/go-build \
    --mount=type=secret,id=netrc,target=/root/.netrc \
    make go-build


# ====== Release stage ======
FROM alpine:${ALPINE_VERSION} AS app

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /go/src/github.com/NishimuraTakuya-nt/go-rest-chi/bin/go-rest-chi ./bin/

CMD ["/app/bin/go-rest-chi"]
