FROM golang:1.25-bookworm AS builder
WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    git \
    openssh-client \
    && rm -rf /var/lib/apt/lists/*

# uncomment for private repos
# ENV GOPRIVATE=XXXXXXX
# RUN git config --global url."ssh://git@github.com/".insteadOf "https://github.com/"

# RUN mkdir -p /root/.ssh && \
#     chmod 700 /root/.ssh && \
#     ssh-keyscan github.com >> /root/.ssh/known_hosts

COPY go.mod go.sum ./

RUN --mount=type=ssh \
    ssh-add -l && \
    go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth-server ./cmd/server


FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /usr/local/bin/server

EXPOSE 9999

ENV SERVER_CONFIG_PATH=/app/config.yaml

WORKDIR /usr/local/bin
CMD ["server"]