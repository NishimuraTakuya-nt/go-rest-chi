# syntax = docker/dockerfile:1

ARG ALPINE_VERSION=3.20
ARG GOLANG_VERSION=1.23.3

# ====== Base stage ======
FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} as base

RUN apk --no-cache add \
    make \
    git \
    gcc \
    musl-dev \
    ca-certificates

WORKDIR /go/src/github.com/NishimuraTakuya-nt/go-rest-chi


# ====== Test stage ======
FROM base as test

COPY go.mod go.sum Makefile ./
RUN --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    make go-download

COPY . .
RUN --mount=type=cache,id=go-test,target=/root/.cache/go-build \
    make test
# change ginkgo command
# id このままで良いのか？ modでいいのか？


# ====== Build stage ======
FROM base as builder

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
FROM alpine:${ALPINE_VERSION} as app

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /go/src/github.com/NishimuraTakuya-nt/go-rest-chi/bin/go-rest-chi ./bin/

CMD ["/app/bin/go-rest-chi"]
