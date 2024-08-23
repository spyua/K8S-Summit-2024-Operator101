# Build the manager binary
FROM golang:1.20-alpine as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
RUN go mod tidy

COPY . .
# COPY main.go main.go
# COPY reconcier.go reconcier.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o operator main.go reconcier.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/operator .
USER 65532:65532

ENTRYPOINT ["/operator"]
