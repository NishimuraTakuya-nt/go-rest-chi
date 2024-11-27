# syntax = docker/dockerfile:1

ARG ALPINE_VERSION=3.20
ARG GOLANG_VERSION=1.23.3
ARG GINKGO_VERSION=v2.22.0
ARG GOLANGCI_LINT_VERSION=v1.62.2

# ====== Base stage ======
FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS base

ARG GINKGO_VERSION
ARG GOLANGCI_LINT_VERSION

RUN apk --no-cache add \
    ca-certificates \
    gcc \
    git \
    make \
    musl-dev \
    tzdata

RUN go install github.com/onsi/ginkgo/v2/ginkgo@${GINKGO_VERSION}
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

WORKDIR /go/src/github.com/NishimuraTakuya-nt/go-rest-chi


# ====== Lint stage ======
FROM base AS lint

COPY go.mod go.sum ./
RUN --mount=type=cache,id=go-deps,target=/go/pkg/mod \
    go mod download

COPY . .
RUN --mount=type=cache,id=go-deps,target=/go/pkg/mod \
    --mount=type=cache,id=go-lint,target=/root/.cache/golangci-lint \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    make lint


# ====== Lint execution stage ======
FROM base AS lint-executor

COPY go.mod go.sum ./
RUN --mount=type=cache,id=go-deps,target=/go/pkg/mod \
    go mod download

COPY . .
CMD ["make", "lint"]


# ====== Test stage ======
FROM base AS test

COPY go.mod go.sum Makefile ./
RUN --mount=type=cache,id=go-deps,target=/go/pkg/mod \
    make go-download

COPY . .
RUN --mount=type=cache,id=go-deps,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    make test-ginkgo-coverage


# ====== Coverage stage ======
FROM scratch AS coverage
COPY --from=test /go/src/github.com/NishimuraTakuya-nt/go-rest-chi/coverage/coverage.txt /coverage.txt


# ====== Test stage for compose ======
FROM base AS test-executor

COPY go.mod go.sum Makefile ./
RUN --mount=type=cache,id=go-deps,target=/go/pkg/mod \
    make go-download

COPY . .
CMD ["make", "test-ginkgo-coverage"]


# ====== Build stage ======
FROM base AS builder

COPY go.mod go.sum Makefile ./
RUN --mount=type=cache,id=go-deps,target=/go/pkg/mod \
    make go-download

COPY . .
RUN --mount=type=cache,id=go-deps,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    make go-build


# ====== Release stage ======
FROM alpine:${ALPINE_VERSION} AS app

RUN apk --no-cache add \
    ca-certificates \
    tzdata

WORKDIR /app

COPY --from=builder /go/src/github.com/NishimuraTakuya-nt/go-rest-chi/bin/go-rest-chi ./bin/

USER nobody:nobody

CMD ["/app/bin/go-rest-chi"]
