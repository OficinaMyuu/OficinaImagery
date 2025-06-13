FROM golang:1.24.4-alpine3.22 AS build

WORKDIR /app

COPY cmd/go.mod cmd/go.sum ./

RUN go mod download

COPY cmd/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./api/main.go

# Final/Run Stage
FROM mcr.microsoft.com/playwright:latest

# Install only Chromium/Chrome
ENV PLAYWRIGHT_BROWSERS_PATH=0
RUN playwright install chrome

RUN apt-get update && \
    apt-get install -y ffmpeg curl && \
    curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /usr/local/bin/yt-dlp && \
    chmod a+rx /usr/local/bin/yt-dlp

WORKDIR /app
COPY --from=build /app/server /app/server

ENV PATH="/app:${PATH}"
ENV TMPDIR="/tmp"

EXPOSE 8080

ENTRYPOINT ["./server"]