# ---- stage 1: Builder ----
FROM golang:1.23-alpine AS builder

WORKDIR /app

# copy module files
COPY go.mod go.sum ./
RUN go mod download

# copy the rest of the application code
COPY . .

# Build the static Go Binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

# ---- stage 2 : runner ----
FROM gcr.io/distroless/base-debian12 AS runner

WORKDIR /app

# copy the complied Binary file from the builder stage
COPY --from=builder /app/server /app/server

# run as non-root user for security
USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/app/server"]

HEALTHCHECK --interval=10s --timeout=2s --retries=3 \

CMD ["/app/server", "health"]
