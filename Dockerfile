ARG REGISTRY

FROM ${REGISTRY}${REGISTRY:+/}golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -mod=vendor -o Mastermind ./cmd/Mastermind

FROM ${REGISTRY}${REGISTRY:+/}alpine
WORKDIR /app
COPY --from=builder /app/Mastermind /bin/
CMD Mastermind
