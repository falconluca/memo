FROM golang:1.17-alpine
WORKDIR /app

ARG TARGETARCH
ARG TARGETOS

ENV GOPROXY=https://goproxy.cn
ENV GO111MODULE=on
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY api/ api/
COPY cmd/ cmd/
COPY internal/ internal/
RUN CGO_ENABLED=0 GOARCH=$TARGETARCH GOOS=$TARGETOS go build -a -o server ./cmd/server/

COPY config/ config/
EXPOSE 50051
ENTRYPOINT ["./server"]