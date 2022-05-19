FROM golang:1.17-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .
ENV GIN_MODE=release
RUN go build

EXPOSE 8080
CMD [ "./meli-challenge-api" ]
