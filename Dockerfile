FROM cgr.dev/chainguard/go:1.20 as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o nungwi ./cmd/nungwi

FROM cgr.dev/chainguard/static

EXPOSE 8080
COPY --from=builder /app/nungwi /nungwi

ENTRYPOINT ["/nungwi"]
