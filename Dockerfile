FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64

WORKDIR /app

COPY . .

RUN go get -d -v ./... \
    && go install -v ./... \
    && go build -o sweater_app ./cmd/main.go

# Export necessary port
EXPOSE 8080

CMD  ["./sweater_app"]