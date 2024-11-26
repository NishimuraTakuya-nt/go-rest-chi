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
RUN --mount=type=cache,id=go-deps,target=/go/pkg/mod \
    make go-download

COPY . .
RUN --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    make test-ginkgo-coverage


# ====== Coverage stage ======
FROM scratch AS coverage
COPY --from=test /go/src/github.com/NishimuraTakuya-nt/go-rest-chi/coverage/coverage.txt /coverage.txt


# ====== Build stage ======
FROM base AS builder

COPY go.mod go.sum Makefile ./
RUN --mount=type=cache,id=go-deps,target=/go/pkg/mod \
    --mount=type=secret,id=netrc,target=/root/.netrc \
    make go-download

COPY . .
RUN --mount=type=cache,id=go-deps,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    --mount=type=secret,id=netrc,target=/root/.netrc \
    make go-build


# ====== Release stage ======
FROM alpine:${ALPINE_VERSION} AS app

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /go/src/github.com/NishimuraTakuya-nt/go-rest-chi/bin/go-rest-chi ./bin/

CMD ["/app/bin/go-rest-chi"]
