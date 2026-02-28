FROM golang AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=nomsgpack -ldflags="-s -w" -o server


FROM scratch

WORKDIR /app

COPY templates/ ./templates/
COPY static/ ./static/
COPY --from=builder /app/server .

EXPOSE 8080
ENTRYPOINT ["./server"]
