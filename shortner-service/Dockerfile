FROM --platform=$BUILDPLATFORM golang:1.26.2

WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY . .

RUN go build -o shortener-service .

EXPOSE 8080

CMD ["./shortener-service"]