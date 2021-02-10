# Build the manager binary
FROM golang:1.13 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY apis/ apis/
COPY apis-products/ apis-products/
COPY controllers/ controllers/
COPY pkg/ pkg/
COPY version/ version/
COPY test/ test/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o rhmi-operator main.go

FROM registry.access.redhat.com/ubi8/ubi:8.0

ENV OPERATOR=/usr/local/bin/rhmi-operator \
    USER_UID=1001 \
    USER_NAME=integreatly-operator

COPY --from=builder /workspace/rhmi-operator /usr/local/bin/rhmi-operator

COPY templates /templates

COPY manifests /manifests

COPY build/bin /usr/local/bin
RUN /usr/local/bin/user_setup


ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
